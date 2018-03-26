package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/headzoo/surf"
	"github.com/jmoiron/jsonq"
)

/**
 *  https://gist.github.com/vool/bbd64eeee313d27a82ab  <== Ryanair
 */
func main() {

	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")

	err := bow.Open("https://api.ryanair.com/farefinder/3/oneWayFares?&departureAirportIataCode=BCN&arrivalAirportIataCode=LPL" +
		"&language=en&limit=16&market=en-gb&offset=0&outboundDepartureDateFrom=2018-05-11&outboundDepartureDateTo=2018-05-12&priceValueTo=150")

	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{}
	body := strings.Replace(bow.Body(), "&#34;", "\"", -1)

	dec := json.NewDecoder(strings.NewReader(body))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	exampleElement, _ := jq.Array("fares")
	fmt.Println("fares", exampleElement)

}
