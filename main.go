package main

import "myurl.com/go/cryptomasters/api"

func main() {
	// Get the rate for the BTC/USD cryptocurrency pair
	rate, err := api.GetRate("BTC")

	// Print the rate (a returned pointer to the Rate data type from the GetRate function above)
	print(rate, err)
}
