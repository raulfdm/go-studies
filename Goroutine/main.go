package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go doSomethingAsync(channel)

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func doSomethingAsync(channel chan string) {
	time.Sleep(2 * time.Second)
	fmt.Printf("2 seconds passed\n")
	channel <- "Done!"
	channel <- "Really..."
	channel <- "I mean... really..."
	channel <- "Done!"
}
