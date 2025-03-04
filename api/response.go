package api

/* You can quickly generate Go structs from JSON
(in this case the CEX crypto API's JSON response here:
https://cex.io/api/ticker/BTC/USD)
using and online JSON to Go struct converter like:
https://mholt.github.io/json-to-go/

There are plenty of online tools/converters like this out there - just Google for one.

Also note the string type metadata on this struct (at the end of each struct property).
These were automatically created by the JSON to Go struct converter above.

This type of struct metadata is generally used for json to Go struct property mapping,
commenting, documentation, or reflection API's that query the structure of your code/data. */

type CEXResponse struct {
	Timestamp             string  `json:"timestamp"`
	Low                   string  `json:"low"`
	High                  string  `json:"high"`
	Last                  string  `json:"last"`
	Volume                string  `json:"volume"`
	Volume30D             string  `json:"volume30d"`
	Bid                   float64 `json:"bid"`
	Ask                   float64 `json:"ask"`
	PriceChange           string  `json:"priceChange"`
	PriceChangePercentage string  `json:"priceChangePercentage"`
	Pair                  string  `json:"pair"`
}
