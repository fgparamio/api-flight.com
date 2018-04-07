package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"../../core/util"

	"github.com/jmoiron/jsonq"
	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")

	util.CheckError(bow.Open("https://www.vivaair.com/pe"))

	// Outputs: "reddit: the front page of the internet"

	fm, _ := bow.Form("form#box-flight-form")
	fm.Set("TravelMode", "OneWay")
	fm.Set("DepartureCity", "LIM")
	fm.Set("ArrivalCity", "CUZ")
	fm.Set("DepartureDate", "2018-03-18")
	fm.Set("ReturnDate", "2018-03-25")
	fm.Set("Currency", "PEN")
	fm.Set("Adults", "1")
	fm.Set("Children", "0")
	fm.Set("Infants", "0")
	fm.Set("DepartureDateForDisplay", "18/04/2018")
	fm.Set("DepartureDateString", "2018-04-18")
	fm.Set("ReturnDateForDisplay", "25/04/2018")
	fm.Set("ReturnDateString", "2018-04-25")
	fm.Set("ExactMonthDate", "exact")

	util.CheckError(fm.Submit())

	var jsonReq = []byte(`{"AvailType":"","searchType":"Week","Currency":"PEN","adults":"1","children":"0","infants":"0","fareDesignator":"",
		"getFromState":false,"isChangeBooking":false,"minimumTotalAmountOut":null,"minimumTotalAmountIn":null,"SelectedPax":null,"culture":"en-US",
		"Flights":[{"DepartureAirport":"LIM","ArrivalAirport":"CUZ","DepartureDate":"2018-04-18T05:00:00.000Z","DepartureDateString":"2018-04-18"},
		{"DepartureAirport":"CUZ","ArrivalAirport":"LIM","DepartureDate":"2018-04-25T05:00:00.000Z","DepartureDateString":"2018-04-25"}]}`)

	bow.Post("https://www.vivaair.com/AsyncAvailability/SearchAsync", "application/json; charset=UTF-8", strings.NewReader(string(jsonReq)))

	data := map[string]interface{}{}
	body := strings.Replace(bow.Body(), "&#34;", "\"", -1)
	dec := json.NewDecoder(strings.NewReader(body))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	exampleElement, _ := jq.Object("availabilityTemplate", "OutAvailability", "0", "FareData")
	fmt.Println("FareData", exampleElement)

}
