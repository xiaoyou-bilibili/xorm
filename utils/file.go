package utils

import (
	"bytes"
	"fmt"
	"os"
)

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// WriteGoFile 文件写入
func WriteGoFile(path, name string, content *bytes.Buffer) error {
	return os.WriteFile(fmt.Sprintf("%s/%s.go", path, name), content.Bytes(), 0775)
}
