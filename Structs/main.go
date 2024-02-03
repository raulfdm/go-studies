package main

import (
	"fmt"
	"time"
)

type User struct {
	name        string
	yearOfBirth int
}

func (user User) CalculateAge() int {
	return time.Now().Year() - user.yearOfBirth
}

func main() {
	user := User{name: "John", yearOfBirth: 1990}

	fmt.Printf("User %s, born in %d, is probably %d years old", user.name, user.yearOfBirth, user.CalculateAge())
}
