package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

// TODO -> This Scraper Not Working. COPA Detect Boot
func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")

	checkError(bow.Open("https://www.copaair.com/es/web"))
	authToken := getAuthToken(bow.Body())
	fmt.Println("AuthToken", authToken)
	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	bow.AddRequestHeader("Cache-Control", "maz-age=0")
	bow.AddRequestHeader("Connection", "keep-alive")
	bow.AddRequestHeader("Content-Type", "application/x-www-form-urlencoded")

	// Outputs: "reddit: the front page of the internet"
	fm, _ := bow.Form("form.row-fluid")

	fm.Set("paymentType", "Money")
	fm.Set("originSearchString", "Bogotá, CO BOG")
	fm.Set("destinationSearchString", "Nueva York, US JFK")
	fm.Set("pos", "CMCO")
	fm.Set("lang", "es")
	fm.Set("outboundOption.originLocationCode:", "BOG")
	fm.Set("outboundOption.destinationLocationCode:", "JFK")
	fm.Set("inboundOption.originLocationCode", "JFK")
	fm.Set("inboundOption.destinationLocationCode", "BOG")
	fm.Set("dateRangeFrom", "abril 28, 2018")
	fm.Set("dateRangeTo", "2018-03-18")
	fm.Set("outboundOption.departureDay", "31")
	fm.Set("outboundOption.departureMonth", "3")
	fm.Set("outboundOption.departureYear", "2018")
	fm.Set("inboundOption.departureDay", "28")
	fm.Set("inboundOption.departureMonth", "4")
	fm.Set("outboundOption.departureYear", "2018")
	fm.Set("tripType", "RT")
	fm.Set("cabinClass", "Economy")
	fm.Set("flexibleSearch", "false")
	fm.Set("guestTypes[0].amount", "1")
	fm.Set("guestTypes[0].type", "ADT")
	fm.Set("guestTypes[1].amount", "0")
	fm.Set("guestTypes[1].type", "CNN")
	fm.Set("guestTypes[2].amount", "0")
	fm.Set("guestTypes[2].type", "INF")
	fm.Set("coupon", "")

	// bow.Open("https://www.copaair.com/api/jsonws/copa-flight-booking-portlet.airport/get-available-destinations/airport-origin-code/BOG/store-front/CO/culture/es-CO?p_auth=" + authToken)
	// bow.Open("https://www.copaair.com/api/jsonws/copa-flight-booking-portlet.route/get-route-operation-date/origin-code/BOG/destination-code/JFK?p_auth=" + authToken)
	// bow.Open("https://www.copaair.com/api/jsonws/copa-flight-booking-portlet.route/get-route-operation-date/origin-code/BOG/destination-code/JFK?p_auth=" + authToken)
	bow.Open("https://www.copaair.com/api/jsonws/copa-flight-booking-portlet.route/get-route-operation-date/origin-code/BOG/destination-code/JFK?p_auth=" + authToken)

	checkError(fm.Submit())

	fmt.Println(bow.Body())

	// form := url.Values{}
	// form.Add("ajaxAction", "true")
	// bow.Post("https://bookings.copaair.com/CMGS/AirLowFareSearchExt.do", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))

	bow.Open("https://bookings.copaair.com/CMGS/AirFareFamiliesForward.do")

	bow.Dom().Find("td.colAirports").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})
}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}

func getAuthToken(body string) string {
	position := strings.Index(body, "ng-init=\"p_au=")
	token := body[position+19 : position+27]
	fmt.Println(token)
	return strings.Replace(token, "\"", "", -1)
}
