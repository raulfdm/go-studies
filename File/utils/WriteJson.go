package file_utils

import "os"

func WriteJson(filename string, data string) error {
	// 0666 is the permission
	// []byte(data) is the content
	return os.WriteFile(filename, []byte(data), 0666)
}
