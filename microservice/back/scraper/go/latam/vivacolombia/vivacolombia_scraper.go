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

	err := bow.Open("https://www.vivacolombia.co")
	if err != nil {
		panic(err)
	}

	// Outputs: "reddit: the front page of the internet"
	fm, _ := bow.Form("form#box-flight-form")
	fm.Set("TravelMode", "OneWay")
	fm.Set("DepartureCity", "BOG")
	fm.Set("ArrivalCity", "MDE")
	fm.Set("DepartureDate", "2018-03-18")
	fm.Set("ReturnDate", "2018-03-25")
	fm.Set("Currency", "COP")
	fm.Set("Adults", "1")
	fm.Set("Children", "0")
	fm.Set("Infants", "0")
	fm.Set("DepartureDateForDisplay", "18/03/2018")
	fm.Set("DepartureDateString", "2018-03-18")
	fm.Set("ReturnDateForDisplay", "25/03/2018")
	fm.Set("ReturnDateString", "2018-04-12")
	fm.Set("ExactMonthDate", "exact")

	checkError(fm.Submit())

	var jsonReq = []byte(`{"AvailType":"","searchType":"Week","Currency":"COP","adults":"1","children":"0","infants":"0","fareDesignator":"",
		"getFromState":false,"isChangeBooking":false,"minimumTotalAmountOut":null,"minimumTotalAmountIn":null,"SelectedPax":null,"culture":"en-US",
		"Flights":[{"DepartureAirport":"BOG","ArrivalAirport":"MDE","DepartureDate":"2018-03-16T05:00:00.000Z","DepartureDateString":"2018-03-16"},
		{"DepartureAirport":"MDE","ArrivalAirport":"BOG","DepartureDate":"2018-03-25T05:00:00.000Z","DepartureDateString":"2018-03-25"}]}`)

	bow.Post("https://vivacolombia.co/AsyncAvailability/SearchAsync", "application/json; charset=UTF-8", strings.NewReader(string(jsonReq)))

	data := map[string]interface{}{}
	body := strings.Replace(bow.Body(), "&#34;", "\"", -1)
	dec := json.NewDecoder(strings.NewReader(body))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	exampleElement, _ := jq.Object("availabilityTemplate", "OutAvailability", "0", "FareData")
	fmt.Println("FareData", exampleElement)
	fmt.Println(body)

}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
