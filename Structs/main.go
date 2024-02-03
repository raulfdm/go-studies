package main

import (
	"fmt"

	"studies.com/structs/models"
)

func main() {

	user := models.User{Name: "John", YearOfBirth: 1990}

	fmt.Println(user)

	animalsThatSpeak := [2]models.Speakable{}

	h1 := models.Human{}
	h1.Name = "John Doe"
	animalsThatSpeak[0] = h1

	parrot := models.Parrot{}
	parrot.Name = "Polly"
	animalsThatSpeak[1] = parrot

	for _, animal := range animalsThatSpeak {
		fmt.Println(animal.Speak("Hello"))
	}
}
