package fileoperate

import (
	"fmt"
	"io"
	"os"
	"path"
)

func CopyFile(src, dst string) error {
	var err error
	var srcFile *os.File
	var destFile *os.File
	var srcInfo os.FileInfo
	if srcFile, err = os.Open(src); err != nil {
		return err
	}
	defer srcFile.Close()
	// dst文件可以不存在，但是上层文件夹必须存在
	if destFile, err = os.Create(dst); err != nil {
		return err
	}
	defer destFile.Close()
	// 拷贝文件内容
	if _, err = io.Copy(destFile, srcFile); err != nil {
		return err
	}
	// 拷贝文件权限

	if srcInfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcInfo.Mode())
}

func CopyDir(src, dst string) error {
	var err error
	var fileList []os.DirEntry
	var srcInfo os.FileInfo
	// 先看看源文件夹是否存在
	if srcInfo, err = os.Stat(src); err != nil {
		return err
	}
	// 参考源文件夹权限，创建目标文件夹
	if err = os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}
	//读取目录下的内容，获取全部文件/文件夹fileinfo切片
	if fileList, err = os.ReadDir(src); err != nil {
		return err
	}
	// 递归拷贝文件
	for _, file := range fileList {
		srcFilePath := path.Join(src, file.Name())
		dstFilePath := path.Join(dst, file.Name())
		if file.IsDir() {
			// 当前遍历到的是一个文件夹 递归
			if err = CopyDir(srcFilePath, dstFilePath); err != nil {
				fmt.Println(err)
			}
		} else {
			if err := CopyFile(srcFilePath, dstFilePath); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
