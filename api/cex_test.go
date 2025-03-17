package api_test

import (
	"testing"

	"myurl.com/go/cryptomasters/api"
)

// A simple, vanilla Go test to test our GetRate function in the cex.go file
func TestAPICall(t *testing.T) {
	/* Passing a blank cryptocurrency ticker to our GetRate function should fail.
	See the related "if len(currency) < 3" code at the top of the GetRate function where we check and return an error for this scenario. */
	_, err := api.GetRate("")

	if err != nil {
		t.Error("The test failed (an error occurred).")
	} else {
		t.Log("The test passed.")
	}

	// Passing a valid three-character or more cryptocurrency ticker to our GetRate function should pass.
	_, err = api.GetRate("BTC")

	if err != nil {
		t.Error("The test failed (an error occurred).")
	} else {
		t.Log("The test passed.")
	}
}
