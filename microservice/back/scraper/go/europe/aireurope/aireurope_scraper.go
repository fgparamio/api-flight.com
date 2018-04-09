package main

import (
	"fmt"

	"../../core/util"

	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")
	util.CheckError(bow.Open("https://www.aireuropa.com/es/vuelos"))

	bow.AddRequestHeader("Accept", "application/json, text/javascript, */*; q=0.01")
	bow.AddRequestHeader("Authorization", "Bearer 8ca33423-2a82-4170-92d5-fcd043dcd83e")
	bow.AddRequestHeader("Content-Type", "application/json;charset=UTF-8")
	bow.AddRequestHeader("Referer", "https://www.aireuropa.com/es/vuelos/vuelos")
	bow.AddRequestHeader("Host", "www.aireuropa.com")
	bow.AddRequestHeader("market", "CO")

	util.CheckError(bow.Open("https://www.aireuropa.com/aireuropaavailability/rest/v3_0/checkout/availability?locale=ES&marketCode=CO" +
		"&application=W&operatingSystem=WEB&versionNumber=1&paxAdult=1&paxChild=0&paxInfant=0&paxAdultResident=0&paxChildResident=0" +
		"&paxInfantResident=0&airportDeparture=MAD&airportArrival=BCN&dateDeparture=22%2F04%2F2018&dateArrival=29%2F04%2F2018&tryType=RT" +
		"&cabineType=&currencyCode=COP&country=CO&idClient=aeaweb&calendarRequest=true&_=" + util.StrTimestampMilli() + "000"))

	jq := util.ToJSONQ(bow.Body())

	exampleElement, _ := jq.ArrayOfObjects("body", "data", "journeys")
	fmt.Println("Availability", exampleElement)

}
