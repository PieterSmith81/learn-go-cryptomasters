package main

import (
	"fmt"

	"myurl.com/go/cryptomasters/api"
)

func main() {
	// Get the rate for the BTC/USD cryptocurrency pair
	rate, err := api.GetRate("BTC")

	// Print the price (last bid price of the BTC/USD cryptocurrency pair and the CEX crypto exchange) to the console
	if err == nil {
		fmt.Printf("The rate for %v is %.2f", rate.Currency, rate.Price)
	} else {
		fmt.Printf("An error occurred. Error details: %v", err)
	}
}
