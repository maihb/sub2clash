package common

import (
	"io"
	"os"
)

func LoadTemplate(templatePath string) ([]byte, error) {
	if _, err := os.Stat(templatePath); err == nil {
		file, err := os.Open(templatePath)
		if err != nil {
			return nil, err
		}
		defer func(file *os.File) {
			if file != nil {
				_ = file.Close()
			}
		}(file)
		result, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, NewFileNotFoundError(templatePath)
}
