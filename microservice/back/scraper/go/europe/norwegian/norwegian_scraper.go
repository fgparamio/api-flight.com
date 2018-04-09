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
	util.CheckError(bow.Open("https://www.norwegian.com/es/"))

	util.CheckError(bow.Open("https://www.norwegian.com/es/reserva/reserve-su-vuelo/preciosbaratos/?A_City=BGO&AdultCount=1&ChildCount=0" +
		"&CurrencyCode=EUR&D_City=OSL&D_Day=28&D_Month=201804&D_SelectedDay=28&IncludeTransit=true&InfantCount=0&R_Day=19&R_Month=201805" +
		"&R_SelectedDay=19&TripType=2&mode=ab"))

	bow.Dom().Find("div.fareCalDayDirectSelected").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

	util.CheckError(bow.Open("https://www.norwegian.com/es/ipc/availability/avaday?D_City=OSL&A_City=BGO&D_SelectedDay=28&D_Day=28" +
		"&D_Month=201804&R_SelectedDay=19&R_Day=19&R_Month=201805&dFare=84&rFare=79&AgreementCodeFK=-1&CurrencyCode=EUR&mode=ab"))

	bow.Dom().Find("td.standardlowfare").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})
}
