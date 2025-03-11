package main

import (
	"fmt"
	"sync"

	"myurl.com/go/cryptomasters/api"
)

func main() {
	currencies := []string{"BTC", "ETH", "SOL"}

	// A WaitGroup is used here to wait for the collection of goroutines anonymous functions below to finish
	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)

		/* An anonymous function that is invoked as a goroutine (that also sets the WaitGroup as done once it's finished).
		Note that these anonymous functions/goroutines have access to the variables in their outer code block, for example, to the currency variable. */
		go func() {
			getCurrencyData(currency)
			wg.Done()
		}()

		/*
			Note that prior to Go version 1.22, there was a loop reference issue which is detailed here:
			https://go.dev/blog/loopvar-preview

			So, prior to Go 1.22, the code above would need to look like the code below to have worked.
			But this common issue has now been fixed in Go 1.22 and above.

			// Pre Go 1.22 code (using closures with concurrency, a.k.a. with goroutines)
			go func(currencyCode string) {
				getCurrencyData(currencyCode)
				wg.Done()
			}(currency)
		*/
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
}
