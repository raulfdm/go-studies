package main

import (
	"fmt"
	"os"

	file_utils "studies.com/file/utils"
)

func main() {
	rootPath, _ := os.Getwd()

	filePath := rootPath + "/data/file.json"

	content, _ := file_utils.ReadJson(filePath)
	fmt.Println(content)

	filePath = rootPath + "/data/file2.json"

	file_utils.WriteJson(filePath, "{\"name\": \"John\"}")
}
