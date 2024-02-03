package main

import (
	"fmt"
	"sync"

	// "sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for count := 0; count < 5; count++ {
		wg.Add(1)

		go func(lambdaCount int) {
			doSomethingAsync(lambdaCount)
			wg.Done()
		}(count)
	}

	wg.Wait()
}

func doSomethingAsync(val int) {
	time.Sleep(2 * time.Second)
	fmt.Printf("2 seconds passed. Val %v\n", val)
}

// Example with channels

// func main() {
// 	channel := make(chan string)

// 	go doSomethingAsync(channel)

// 	fmt.Println(<-channel)
// 	fmt.Println(<-channel)
// 	fmt.Println(<-channel)
// 	fmt.Println(<-channel)
// }

// func doSomethingAsync(channel chan string) {
// 	time.Sleep(2 * time.Second)
// 	fmt.Printf("2 seconds passed\n")
// 	channel <- "Done!"
// 	channel <- "Really..."
// 	channel <- "I mean... really..."
// 	channel <- "Done!"
// }
