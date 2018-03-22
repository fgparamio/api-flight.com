package main

import (
	"fmt"

	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()

	checkError(bow.Open("https://bookings.sita.aero/itd/itd/eq"))

	bow.Open("https://bookings.sita.aero/itd/itd/eq/lang/es/travel/quotes?journeyCount=2&search=JOURNEY" +
		"&departDate0=2018-03-30&departLocation0=Airport.UIO&arriveLocation0=Airport.CUE&departDate1=2018-04-06" +
		"&departLocation1=Airport.CUE&arriveLocation1=Airport.UIO&numInfants=0&numAdults=1&numChildren=0" +
		"&fareClass=Economy&concessionaryTravellers=false&trAttr_ecuadorianOrResident=false&step=Search")

	bow.Dom().Find("form").AddClass("classForm")
	fm, _ := bow.Form("form.classForm")
	bow.Click("div.robot-code-center")
	fm.Submit()

	bow.Open("https://bookings.sita.aero/itd/itd/eq/lang/es/travel/quotes?journeyCount=2&search=JOURNEY" +
		"&departDate0=2018-03-30&departLocation0=Airport.UIO&arriveLocation0=Airport.CUE&departDate1=2018-04-06" +
		"&departLocation1=Airport.CUE&arriveLocation1=Airport.UIO&numInfants=0&numAdults=1&numChildren=0" +
		"&fareClass=Economy&concessionaryTravellers=false&trAttr_ecuadorianOrResident=false&step=Search")

	fmt.Println(bow.Dom().Html())

}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
