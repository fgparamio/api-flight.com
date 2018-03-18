package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jmoiron/jsonq"
	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")

	err := bow.Open("https://booking.flybondi.com/es/Flight/Search?adults=1&children=0&currency=ARS&departureDateString=2018-04-17" +
		"&fromCityCode=EPA&infants=0&returnDateString=2018-04-28&roundTrip=true&toCityCode=COR&useFlexDates=true")

	// Check GET Request
	checkError(err)

	var jsonReq = []byte(`{"interline":false,"fromCityCode":"EPA","toCityCode":"COR","departureDateString":"2018-04-17",
		"returnDateString":"2018-04-28","startDateStringOutbound":"2018-04-17","endDateStringOutbound":"2018-04-17",
		"startDateStringInbound":"2018-04-28","endDateStringInbound":"2018-04-28","adults":1,"children":0,"infants":0,
		"roundTrip":true,"useFlexDates":true,"isOutbound":true,"filterMethod":"100","promocode":"","currency":"ARS",
		"languageCode":"es-AR","fareTypeCategory":1,"IATANumber":"","securityToken":""}`)

	bow.Post("https://booking.flybondi.com/Api/AvailablityRequest/Post", "application/json; charset=UTF-8", strings.NewReader(string(jsonReq)))

	data := map[string]interface{}{}
	body := strings.Replace(bow.Body(), "&#34;", "\"", -1)
	dec := json.NewDecoder(strings.NewReader(body))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	exampleElement, _ := jq.Object("Availability", "OutboundSegments", "0")
	fmt.Println("First Segment", exampleElement)
}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
