package file

import (
	"os"
)

// 判断路径是否存在IsExist
func IsExist(path string) bool {

	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 创建路径 Mkdir
func Mkdir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// WriteFile 写入文件
func WriteFile(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
