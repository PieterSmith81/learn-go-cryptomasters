package main

import (
	"fmt"
	"sync"

	"myurl.com/go/cryptomasters/api"
)

// A WaitGroup package-scoped variable is used here to wait for the collection of goroutines in main() below to finish
var wg sync.WaitGroup

func main() {
	currencies := []string{"BTC", "ETH", "SOL"}

	for _, currency := range currencies {
		wg.Add(1)
		go getCurrencyData(currency)
	}

	wg.Wait()
}

func getCurrencyData(currency string) {
	// Get the rate for the BTC/USD cryptocurrency pair
	rate, err := api.GetRate(currency)

	// Print the price (last bid price of the BTC/USD cryptocurrency pair and the CEX crypto exchange) to the console
	if err == nil {
		fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price)
	} else {
		fmt.Printf("An error occurred. Error details: %v", err)
	}

	wg.Done()
}
