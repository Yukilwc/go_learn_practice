package fileoperate

import (
	"io"
	"os"
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
	if destFile, err = os.Open(dst); err != nil {
		return err
	}
	defer destFile.Close()
	// 拷贝文件内容
	if _, err = io.Copy(destFile, srcFile); err != nil {
		return nil
	}
	// 拷贝文件权限

	if srcInfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcInfo.Mode())
}
