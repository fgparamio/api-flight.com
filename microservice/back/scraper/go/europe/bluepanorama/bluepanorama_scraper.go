package main

import (
	"fmt"

	"../../core/util"

	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()

	util.CheckError(bow.Open("https://www.blue-panorama.com/main"))

	bow.Open("https://ibe-app.blue-panorama.com/ibe-rest-api/service/special-service/oneday?origin=FCO&destination=HER")

	bow.Open("https://ibe-app.blue-panorama.com/ibe-rest-api/search/flight/calendar/fares?tripType=RT&pax=1,ADULT" +
		"&pax=0,CHILD&pax=0,INFANT&oda=HER&ooa=FCO&od=2018-05-27&oed=2018-06-26&rda=FCO&roa=HER&rd=2018-05-27" +
		"&red=2018-06-26&findNearbyAirports=false")

	jq := util.ToJSONQ(bow.Body())
	trips, _ := jq.Object("trips")
	fares, _ := jq.Object("onwardFares")
	fmt.Println("Fares", fares)
	fmt.Println("Going", trips)

	bow.Open("https://ibe-app.blue-panorama.com/ibe-rest-api/search/flight/calendar/fares?tripType=RT&pax=1,ADULT" +
		"&pax=0,CHILD&pax=0,INFANT&oda=HER&ooa=FCO&od=2018-05-27&oed=2018-06-26&rda=FCO&roa=HER&rd=2018-05-28" +
		"&red=2018-06-28&findNearbyAirports=false&sod=2018-05-27")

	jq = util.ToJSONQ(bow.Body())
	trips, _ = jq.Object("trips")
	fares, _ = jq.Object("returnFares")
	fmt.Println("Fares", fares)
	fmt.Println("Trips", trips)

}
