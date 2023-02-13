package main

import (
	"encoding/json"
	"fmt"
	"go_test_init/fileoperate"
	"os"
	"path"
	"strings"
	"time"
)

var mpFolder = "D:/workspace/work/web/gitee/etranscode/etransmp3IM"
var backupFolder = "D:/workspace/work/web/gitee/etranscode/etransmp3Backup"

func main() {
	fmt.Println("=====小程序同步工具=====")
	fmt.Println("输入quit退出程序")
	fmt.Println("小程序路径:", mpFolder)
	fmt.Println("备份文件夹路径:", backupFolder)
	loadConfig()
	loop()
}

func loadConfig() error {
	content, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println("未找到配置文件，使用默认配置")
		return err
	}
	data := struct {
		MpFolder     string
		BackupFolder string
	}{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Println("JSON解析错误:", err)
		return err
	}
	fmt.Println("JSON解析成功")
	mpFolder = data.MpFolder
	backupFolder = data.BackupFolder
	return nil
}

func loop() {
	for {
		var text string
		fmt.Print("请输入你的转移目标,mp为备份文件夹=>小程序,backup为小程序=>备份文件夹(mp/backup):", text)
		fmt.Scanln(&text)
		if text == "exit" || text == "quit" {
			break
		}
		if text == "backup" {
			err := mp2backup()
			if err != nil {
				fmt.Println("mp2backup失败:", err)
				continue
			}
			fmt.Println("从小程序迁移到备份文件夹完成!")
		} else if text == "mp" {
			err := backup2mp()
			if err != nil {
				fmt.Println("backup2mp失败:", err)
				continue
			}
			fmt.Println("从备份文件夹迁移到小程序完成!")

		} else {
			fmt.Printf("输入的%s选项是无效选项\n", text)
		}
	}
}

// 从小程序迁移到备份
func mp2backup() error {
	var indexName string
	fmt.Print("小程序=>备份文件夹，请输入转移的首页名称(例如:index10):")
	fmt.Scanln(&indexName)
	var err error
	// 拷贝首页代码
	mpIndexPath := path.Join(mpFolder, "pages/index")
	backupIndexPath := path.Join(backupFolder, "pages/_temp_index_"+time.Now().Format("20060102150405"))
	err = fileoperate.CopyDir(mpIndexPath, backupIndexPath)
	if err != nil {
		return err
	}
	fmt.Println("首页代码迁移完成!")
	// 拷贝首页图片
	mpImagesPath := path.Join(mpFolder, "images", indexName)
	backupImagesPath := path.Join(backupFolder, "images/_temp_images_"+time.Now().Format("20060102150405"))
	err = fileoperate.CopyDir(mpImagesPath, backupImagesPath)
	if err != nil {
		return err
	}
	fmt.Println("首页图片迁移完成!")

	// 首页代码清理与重命名
	backupOriginIndexPath := path.Join(backupFolder, "pages", indexName)
	if err = clearAndRename(backupOriginIndexPath, backupIndexPath); err != nil {
		return err
	}

	// 首页图片清理与重命名
	backupOriginImagesPath := path.Join(backupFolder, "images", indexName)
	if err = clearAndRename(backupOriginImagesPath, backupImagesPath); err != nil {
		return err
	}

	return nil
}

func backup2mp() error {
	var indexName string
	fmt.Print("备份文件夹=>小程序，请输入转移的首页名称(例如:index10):")
	fmt.Scanln(&indexName)
	var err error
	// 拷贝首页代码
	mpIndexPath := path.Join(mpFolder, "pages/_temp_index_"+time.Now().Format("20060102150405"))
	backupIndexPath := path.Join(backupFolder, "pages", indexName)
	err = fileoperate.CopyDir(backupIndexPath, mpIndexPath)
	if err != nil {
		return err
	}
	fmt.Println("首页代码迁移完成!")
	// 拷贝首页图片
	mpImagesPath := path.Join(mpFolder, "images", "_temp_images_"+time.Now().Format("20060102150405"))
	backupImagesPath := path.Join(backupFolder, "images", indexName)
	err = fileoperate.CopyDir(backupImagesPath, mpImagesPath)
	if err != nil {
		return err
	}
	fmt.Println("首页图片迁移完成!")
	// 首页代码清理与重命名
	mpOriginIndexPath := path.Join(mpFolder, "pages", "index")
	if err = clearAndRename(mpOriginIndexPath, mpIndexPath); err != nil {
		return err
	}

	// 首页图片清理与重命名
	mpOriginImagesPath := path.Join(mpFolder, "images", indexName)
	if err = clearIndexPrefixFolderAndRename(mpOriginImagesPath, mpImagesPath); err != nil {
		return err
	}
	return nil
}

// 清除之前路径的文件，把临时文件重命名
func clearAndRename(correctPath, tempPath string) error {
	var err error
	if _, err = os.Stat(correctPath); err == nil {
		fmt.Println("备份文件夹中源文件存在,移除源文件:", correctPath)
		if err = os.RemoveAll(correctPath); err != nil {
			return err
		}
	} else if os.IsNotExist(err) {
		fmt.Println("备份文件夹中源文件不存在,则不做处理", correctPath)
	} else {
		return err
	}
	if err = os.Rename(tempPath, correctPath); err != nil {
		return err
	}
	return nil
}

// 清除旧的index开头的图片文件夹，然后重命名临时文件夹
func clearIndexPrefixFolderAndRename(correctPath, tempPath string) error {
	var err error
	var fileList []os.DirEntry
	mpImagesFolder := path.Join(mpFolder, "images")
	if fileList, err = os.ReadDir(mpImagesFolder); err != nil {
		return err
	}
	for _, file := range fileList {
		if file.IsDir() {
			// 是文件夹
			name := file.Name()
			if strings.HasPrefix(name, "index") {
				indexFolder := path.Join(mpImagesFolder, name)
				if err = os.RemoveAll(indexFolder); err != nil {
					return err
				}
				fmt.Println("有index前缀，移除成功:", indexFolder)

			}
		}
	}
	if err = os.Rename(tempPath, correctPath); err != nil {
		return err
	}

	return nil
}
