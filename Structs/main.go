package main

import (
	"fmt"

	"studies.com/structs/models"
)

func main() {

	user := models.User{Name: "John", YearOfBirth: 1990}

	fmt.Println(user)
}
