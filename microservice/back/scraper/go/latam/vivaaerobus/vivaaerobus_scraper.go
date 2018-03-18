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

	err := bow.Open("https://www.vivaaerobus.com/mx")
	if err != nil {
		panic(err)
	}

	// Outputs: "reddit: the front page of the internet"
	fm, _ := bow.Form("form#box-flight-form")
	fm.Set("TravelMode", "1")
	fm.Set("DepartureCity", "QRO")
	fm.Set("choseninput", "")
	fm.Set("ArrivalCity", "CUN")
	fm.Set("Adults", "1")
	fm.Set("Children", "0")
	fm.Set("Infants", "0")
	fm.Set("DepartureDate", "2018-04-04")
	fm.Set("DepartureDateForDisplay", "4 Abr")
	fm.Set("DepartureDateString", "2018-04-04")
	fm.Set("ReturnDate", "2018-05-26")
	fm.Set("ReturnDateForDisplay", "26 May")
	fm.Set("ReturnDateString", "2018-05-26")
	fm.Set("Currency", "MXN")
	fm.Set("PromoCode", "VIVA")
	fm.Set("SpecialCampaignCode", "")

	checkError(fm.Submit())

	var jsonReq = []byte(`{"AvailType":"","searchType":"Week","Currency":"MXN","adults":"1","children":"0",
		"infants":"0","fareDesignator":"VIVA","getFromState":false,"isChangeBooking":false,"minimumTotalAmountOut":null,
		"minimumTotalAmountIn":null,"SelectedPax":null,"culture":"en-US",
		"Flights":[{"DepartureAirport":"QRO","ArrivalAirport":"CUN","DepartureDate":"2018-04-04T05:00:00.000Z","DepartureDateString":"2018-04-04"},
		{"DepartureAirport":"CUN","ArrivalAirport":"QRO","DepartureDate":"2018-05-26T05:00:00.000Z","DepartureDateString":"2018-05-26"}]}`)

	bow.Post("https://www.vivaaerobus.com/AsyncAvailability/SearchAsync", "application/json; charset=UTF-8", strings.NewReader(string(jsonReq)))

	data := map[string]interface{}{}
	body := strings.Replace(bow.Body(), "&#34;", "\"", -1)
	dec := json.NewDecoder(strings.NewReader(body))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	exampleElement, _ := jq.Object("availabilityTemplate", "OutAvailability", "0", "FareData")
	fmt.Println("FareData", exampleElement)

}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
