package user

import "time"

type User struct {
	Name        string
	YearOfBirth int
}

func (user User) CalculateAge() int {
	return time.Now().Year() - user.YearOfBirth
}
