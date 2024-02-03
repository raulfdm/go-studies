package main

import (
	"fmt"

	"studies.com/cryptomasters/api"
)

func main() {
	rate, err := api.GetRates("EUR")

	if err != nil {
		panic(err)
	}

	fmt.Printf("1 BTC is %.2f %v\n", rate.Price, rate.Currency)
}
