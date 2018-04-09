package main

import (
	"fmt"
	"strings"

	"../../core/util"

	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")
	util.CheckError(bow.Open("https://www.aerlingus.com/html/es-US/home.html"))
	util.CheckError(bow.Open("https://www.aerlingus.com/html/flightSearchResult.html"))

	var jsonReq = []byte(`{"flightJourneySearches":[{"sourceAirportCode":"MAD","destinationAirportCode":"DUB","departureDate":"2018-05-09"},
		{"sourceAirportCode":"DUB","destinationAirportCode":"MAD","departureDate":"2018-05-17"}],
		"fareType":"RETURN","fareCategory":"ECONOMY","numAdults":1,"numYoungAdults":0,"numChildren":0,
		"numInfants":0,"promoCode":"","groupBooking":"false"}`)

	bow.AddRequestHeader("Accept", "application/json, text/plain, */*")
	bow.AddRequestHeader("Referer", "https://www.aerlingus.com/html/flightSearchResult.html")
	bow.AddRequestHeader("Host", "www.aerlingus.com")
	bow.AddRequestHeader("Origin", "https://www.aerlingus.com")

	bow.Post("https://www.aerlingus.com/api/search/fixedFlight", "application/json;charset=UTF-8", strings.NewReader(string(jsonReq)))

	jq := util.ToJSONQ(bow.Body())

	exampleElement, _ := jq.ArrayOfObjects("data", "0", "flightOptions")
	fmt.Println("Availability", exampleElement)

}
