package main

import (
	"fmt"
	"go_test_init/fileoperate"
	"os"
	"path"
	"time"
)

func main() {
	fmt.Println("=====小程序同步工具=====")
	fmt.Println("输入quit退出程序")
	fmt.Println("小程序路径:", mpFolder)
	fmt.Println("备份文件夹路径:", backupFolder)
	loop()
}

var mpFolder = "D:/workspace/work/web/gitee/etranscode/etransmp3IM"
var backupFolder = "D:/workspace/work/web/gitee/etranscode/etransmp3Backup"

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
	fmt.Print("请输入转移的首页名称(例如:index10):")
	fmt.Scanln(&indexName)
	var err error
	// 拷贝首页代码
	mpIndexPath := path.Join(mpFolder, "pages/index")
	backupIndexPath := path.Join(backupFolder, "pages/index_temp_"+time.Now().Format("20060102150405"))
	err = fileoperate.CopyDir(mpIndexPath, backupIndexPath)
	if err != nil {
		return err
	}
	fmt.Println("首页代码迁移完成!")
	// 拷贝首页图片
	mpImagesPath := path.Join(mpFolder, "images", indexName)
	backupImagesPath := path.Join(backupFolder, "images/index_temp_"+time.Now().Format("20060102150405"))
	err = fileoperate.CopyDir(mpImagesPath, backupImagesPath)
	if err != nil {
		return err
	}
	fmt.Println("首页图片迁移完成!")

	// 首页代码清理与重命名
	backupOriginIndexPath := path.Join(backupFolder, "pages", indexName)
	if _, err = os.Stat(backupOriginIndexPath); err == nil {
		fmt.Println("备份文件夹中源文件存在,移除源文件:", backupOriginIndexPath)
		if err = os.RemoveAll(backupOriginIndexPath); err != nil {
			return err
		}
	} else if os.IsNotExist(err) {
		fmt.Println("备份文件夹中源文件不存在,则不做处理", backupOriginIndexPath)
	} else {
		return err
	}
	if err = os.Rename(backupIndexPath, backupOriginIndexPath); err != nil {
		return err
	}
	// 首页图片清理与重命名
	backupOriginImagesPath := path.Join(backupFolder, "images", indexName)
	if _, err = os.Stat(backupOriginImagesPath); err == nil {
		fmt.Println("备份文件夹中源文件存在,移除源文件:", backupOriginImagesPath)
		if err = os.RemoveAll(backupOriginImagesPath); err != nil {
			return err
		}
	} else if os.IsNotExist(err) {
		fmt.Println("备份文件夹中源文件不存在,则不做处理", backupOriginImagesPath)
	} else {
		return err
	}
	if err = os.Rename(backupImagesPath, backupOriginImagesPath); err != nil {
		return err
	}

	return nil
}

func backup2mp() error {
	return nil
}
