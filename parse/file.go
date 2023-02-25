package parse

import "os"

// FileExists 文件是否存在
func FileExists(name string) bool {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

// WriteFile 写内容进入文件
func WriteFile(file *os.File, content string) error {
	if _, err := file.Write([]byte(content)); err != nil {
		return err
	}
	return nil
}
