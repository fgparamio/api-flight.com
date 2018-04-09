package main

import (
	"fmt"
	"strings"

	"../../core/util"
	"github.com/PuerkitoBio/goquery"

	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")
	util.CheckError(bow.Open("https://www.alitalia.com/en_en"))

	fm, _ := bow.Form("form#cercaVoliForm")
	fm.Set("Origin", "ROM")
	fm.Set("departureInput", "Rome ROM")
	fm.Set("CUG", "ADT")
	fm.Set("Destination", "MIL")
	fm.Set("arrivalInput", "Milan MIL")
	fm.Set("SearchType", "a+r")
	fm.Set("DepartureDate", "29/04/2018")
	fm.Set("ReturnDate", "04/05/2018")
	fm.Set("Adults", "1")
	fm.Set("Children", "0")
	fm.Set("Infants", "0")
	fm.Set("classOfFlight", "economy")
	fm.Set("_isAjax", "true")
	fm.Set("_action", "validate")

	util.CheckError(fm.Submit())

	bow.Open("https://www.alitalia.com/libs/granite/csrf/token.json")
	bow.Open("https://www.alitalia.com/libs/cq/i18n/dict.en-en.json")
	bow.Open("https://www.alitalia.com/en_en/booking/flight-select.html")
	bow.Open("https://www.alitalia.com/en_en/booking/flight-select.bookingchecknavigationprocessconsumer.json?destPhase=flightsSearch&nocache=" + util.StrLargeTimestampMilli() + "&residency=&_=" + util.StrLargeTimestampMilli())
	bow.Post("https://www.alitalia.com/en_en/home-page/.bookingflightsearchserviceconsumer.json", "application/x-www-form-urlencoded; charset=UTF-8",
		strings.NewReader("isResidency=false&_isAjax=true"))
	bow.Open("https://www.alitalia.com/en_en/booking/flight-select.flight-select-main-partial.html?_=" + util.StrLargeTimestampMilli() + "000")

	fmt.Println(bow.Body())

	bow.Dom().Find("div.bookingTable__body").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})
}
