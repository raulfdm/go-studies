package main

import (
	"fmt"
	"sync"

	"studies.com/cryptomasters/api"
)

func main() {
	currencies := []string{"USD", "EUR"}

	var wg sync.WaitGroup
	defer wg.Wait()

	for _, currency := range currencies {
		wg.Add(1)

		go func(currency string) {
			defer wg.Done()
			getRate(currency)
		}(currency)
	}

}

func getRate(currency string) {
	rate, err := api.GetRates(currency)

	if err != nil {
		panic(err)
	}

	fmt.Printf("1 BTC is %.2f %v\n", rate.Price, rate.Currency)
}
