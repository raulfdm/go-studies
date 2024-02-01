package file_utils

import (
	"os"
)

func ReadJson(filename string) (string, error) {
	content, err := os.ReadFile(filename)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

func WriteJson(filename string, data string) {
	// write data to file
}
