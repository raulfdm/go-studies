package models

type Speakable interface {
	Speak(something string) string
}
