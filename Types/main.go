package main

import "fmt"

type celsius int8
type fahrenheit float32

func main() {
	f := fahrenheit(32)
	c := f.toCelsius()

	fmt.Printf("Celsius: %v\n", f.toCelsius())
	fmt.Printf("Fahrenheit: %v\n", c.toFahrenheit())
}

func (f fahrenheit) toCelsius() celsius {
	return celsius((f - 32) * 5 / 9)
}

func (c celsius) toFahrenheit() fahrenheit {
	return fahrenheit(c*9/5 + 32)
}
