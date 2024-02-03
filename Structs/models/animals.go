package models

type Animal struct {
	Name string
}

type Dog struct {
	Animal
}

type Parrot struct {
	Animal
}

func (parrot Parrot) Speak(something string) string {
	return "I am an Parrot 🦜 and I can speak: " + something
}

type Human struct {
	Animal
}

func (human Human) Speak(something string) string {
	return "I am an Human and I can speak: " + something
}
