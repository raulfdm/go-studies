package main

import (
	"fmt"

	models "studies.com/structs/models"
)

func main() {

	user := models.User{Name: "John", YearOfBirth: 1990}

	fmt.Printf("User %s, born in %d, is probably %d years old", user.Name, user.YearOfBirth, user.CalculateAge())
}
