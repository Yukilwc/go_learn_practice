package utils

import (
	"errors"
	"os"
)

func FolderExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		if fi.IsDir() {
			return true, nil
		} else {
			return false, errors.New("存在同名文件")
		}
	}
}
