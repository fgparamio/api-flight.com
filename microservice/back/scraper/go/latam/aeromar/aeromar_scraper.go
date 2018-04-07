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

	util.CheckError(bow.Open("https://www.aeromar.com.mx/"))

	fm, _ := bow.Form("[name='send_the_data']")
	fm.Set("flight_dep", "")
	fm.Set("flight_rtn", "")
	fm.Set("flight_type", "RT")
	fm.Set("flight_from", "2018-03-29")
	fm.Set("flight_to", "2018-04-05")
	fm.Set("flight_adults", "1")
	fm.Set("flight_kids", "0")
	fm.Set("flight_infants", "")

	fm.Set("lang", "es")
	fm.Set("fill_from", "Cancún (CUN)")
	fm.Set("fill_to", "Mérida (MID)")
	fm.Set("datepicker", "31 Mar, 2018")
	fm.Set("corporate_id", "Ingresa aquí tu código ID")

	util.CheckError(fm.Submit())

	bow.Open("https://wl-prod.havail.sabresonicweb.com/SSW2010/VWVW/webqtrip.html?earchType=NORMAL&journeySpan=RT" +
		"&origin=CUN&destination=MID&departureDate=2018-03-31&returnDate=2018-04-07&numAdults=1" +
		"&numChildren=0&numInfants=0&alternativeLandingPage=true&lang=es_ES")

	bow.Dom().Find("tr.yui-dt-even").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}
