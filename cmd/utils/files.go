package utils

import (
	"fmt"
	"io"
	"os"
)

func CreateOrOpenFileAndRead(path string) ([]byte, error) {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil { return []byte(""), err }
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return []byte(""), fmt.Errorf("fail to read file (%v): %v\n", path, err)
	}

	return content, nil
}

func CreateOrOpenFileAndRewrite(path string, content []byte) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil { return err }
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return fmt.Errorf("fail to write to file (%v): %v\n", path, err)
	}
	return nil
}