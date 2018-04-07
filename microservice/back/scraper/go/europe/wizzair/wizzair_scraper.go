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

	util.CheckError(bow.Open("https://wizzair.com/#/booking/select-flight/SOF/BUD/2018-05-12/2018-05-26/1/0/0"))

	var jsonReq = []byte(`{"isFlightChange":false,"isSeniorOrStudent":false,"flightList":[{"departureStation":"SOF","arrivalStation":"BUD",
		"departureDate":"2018-05-12"},{"departureStation":"BUD","arrivalStation":"SOF",
		"departureDate":"2018-05-26"}],"adultCount":1,"childCount":0,"infantCount":0,
		"wdc":true,"rescueFareCode":""}`)

	bow.Post("https://be.wizzair.com/7.10.1/Api/search/search", "application/json; charset=UTF-8", strings.NewReader(string(jsonReq)))

	jq := util.ToJSONQ(bow.Body())

	exampleElement, _ := jq.Array("outboundFlights")
	fmt.Println("Availability", exampleElement)

}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
