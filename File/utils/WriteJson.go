package file_utils

import "os"

func WriteJson(filename string, data string) error {
	return os.WriteFile(filename, []byte(data), 0666)
}
