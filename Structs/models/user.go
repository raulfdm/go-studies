package models

import (
	"fmt"
	"time"
)

type User struct {
	Name        string
	YearOfBirth int
}

func (user User) CalculateAge() int {
	return time.Now().Year() - user.YearOfBirth
}

func (user User) String() string {
	return fmt.Sprintf("--- %v ----\n - Born in %v", user.Name, user.YearOfBirth)
}
