package file_utils

import (
	"os"
)

func ReadJson(filename string) string {
	content, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(content)
}

func writeJson(filename string, data string) {
	// write data to file
}
