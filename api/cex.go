package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"myurl.com/go/cryptomasters/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	upCurrency := strings.ToUpper(currency)
	// Perform a http get request
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))

	// If an error occurred during the http get request
	if err != nil {
		return nil, err
	}

	var response CEXResponse

	if res.StatusCode == http.StatusOK {
		// If the http get response was OK

		// Read the http body "chunks" (as bytes) over the network/internet, waiting for the entire http body to be "downloaded"
		bodyBytes, err := io.ReadAll(res.Body)

		// If an error occurred while reading the http body "chunks"
		if err != nil {
			return nil, err
		}

		// JSON parsing (with Unmarshal)
		err = json.Unmarshal(bodyBytes, &response)

		if err != nil {
			return nil, err
		}

	} else {
		// If the http get response was NOT OK

		// Return an error using the format package's Errorf function
		return nil, fmt.Errorf("HTTP request error status code received: %v", res.StatusCode)
	}

	// Return the Rate as a pointer to a struct
	rate := datatypes.Rate{Currency: currency, Price: response.Bid}

	return &rate, nil
}
