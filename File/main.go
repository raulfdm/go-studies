package main

import (
	"fmt"

	file_utils "studies.com/file/utils"
)

func main() {
	content := file_utils.ReadJson("./data/file.json")

	fmt.Println(content)
}
