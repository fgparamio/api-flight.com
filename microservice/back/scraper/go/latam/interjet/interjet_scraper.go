package main

import (
	"fmt"

	"../../core/util"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")

	util.CheckError(bow.Open("https://www.interjet.com/es-mx"))

	// Outputs: "reddit: the front page of the internet"
	fm, _ := bow.Form("form#SkySales")
	fm.Set("__EVENTTARGET", "")
	fm.Set("__EVENTARGUMENT", "BOG")
	fm.Set("__VIEWSTATE", "")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$RadioButtonMarketStructure", "RoundTrip")
	fm.Set("OriginStation", "Ciudad de México")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$TextBoxMarketOrigin1", "MEX")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$DropDownListMarketDay1", "4")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$DropDownListMarketMonth1", "2018-04")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$DropDownListMarketDateRange1", "0|0")
	fm.Set("date_picker", "San Diego vía Cross Border Xpress")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$TextBoxMarketDestination1", "TIJ")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$DropDownListMarketDay2", "28")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$DropDownListMarketMonth2", "2018-04")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$DropDownListMarketDateRange2", "0|0")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$DropDownListResidentCountry", "MX")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$DropDownListMarketMonth2", "2018-04")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$DropDownListPassengerType_ADT", "1")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$DropDownListPassengerType_CHD", "0")
	fm.Set("DROPDOWNLISTPASSENGERTYPE_INFANT", "0")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$DropDownListPassengerType_INFANT", "0")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$DropDownListFareTypes", "")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$DropDownListPaxDiscount", "INET")
	fm.Set("SkysalesSearchFlightBy", "on")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$DropDownListSearchBy", "columnView")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$DropDownListFarePreference", "")
	fm.Set("ControlGroupHomeView$AvailabilitySearchInputHomeView$ButtonSubmit", "Buscar vuelos")
	fm.Set("OriginStations", "")
	fm.Set("DestinationStations", "")

	util.CheckError(fm.Submit())
	// Parse Example

	bow.Dom().Find("div.fare-data").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}
