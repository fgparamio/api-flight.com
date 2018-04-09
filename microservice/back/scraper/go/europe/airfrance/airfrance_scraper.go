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
	util.CheckError(bow.Open("https://www.airfrance.fr"))

	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	bow.AddRequestHeader("Referer", "https://www.airfrance.fr")
	bow.AddRequestHeader("Host", "www.airfrance.fr")
	bow.AddRequestHeader("Origin", "https://www.airfrance.fr")
	bow.AddRequestHeader("Content-Type", "application/x-www-form-urlencoded")

	fm, _ := bow.Form("form#minibe__form")
	fm.Set("typeTrip", "2")
	fm.Set("departure", "PAR")
	fm.Set("departureType", "CITY")
	fm.Set("departureType", "AIRP")
	fm.Set("arrival", "BOD")
	fm.Set("arrivalType", "AIRP")
	fm.Set("arrivalType", "CITY")
	fm.Set("dayDate", "26")
	fm.Set("yearMonthDate", "201804")
	fm.Set("dayDate", "30")
	fm.Set("yearMonthDate", "201804")
	fm.Set("nbPassenger", "1")
	fm.Set("paxTypoList", "ADT")
	fm.Set("selectCabin", "on")
	fm.Set("plusOptions", "")
	fm.Set("nbAdults", "1")
	fm.Set("nbChildren", "0")
	fm.Set("nbEnfants", "0")
	fm.Set("nbBebes", "0")
	fm.Set("cabin", "Y")
	fm.Set("subCabin", "MCHER")
	fm.Set("haul", "SH")
	fm.Set("familyTrip", "NON")
	fm.Set("calendarSearch", "1")
	fm.Set("flyingBlueMember", "false")
	fm.Set("partnerRequest", "false")
	fm.Set("corporateMode", "false")
	fm.Set("isUM", "false")
	fm.Set("optionalUM", "false")
	fm.Set("mandatoryUM", "true")
	fm.Set("standardMandatory", "true")
	fm.Set("subscriberHOPContext", "false")
	fm.Set("notFromFlight", "true")

	util.CheckError(fm.Submit())

	var jsonReq = []byte(`typeTrip=2&departure=PAR&departure=BOD&departureType=CITY&departureType=AIRP&arrival=BOD&arrival=PAR
		&arrivalType=AIRP&arrivalType=CITY&dayDate=26&dayDate=30&yearMonthDate=201804&yearMonthDate=201804&nbPassenger=1
		&paxTypoList=ADT&selectCabin=on&plusOptions=&nbAdults=1&nbChildren=0&nbEnfants=0&nbBebes=0&cabin=Y&subCabin=MCHER
		&haul=SH&familyTrip=NON&calendarSearch=1&flyingBlueMember=false&partnerRequest=false&corporateMode=false&isUM=false
		&optionalUM=false&mandatoryUM=true&standardMandatory=true&subscriberHOPContext=false&notFromFlight=true`)

	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	bow.AddRequestHeader("Referer", "https://www.airfrance.fr/FR/fr/local/process/standardbooking/ValidateSearchAction.do")
	bow.AddRequestHeader("Host", "www.airfrance.fr")
	bow.AddRequestHeader("Origin", "https://www.airfrance.fr")
	bow.AddRequestHeader("Content-Type", "application/x-www-form-urlencoded")
	bow.AddRequestHeader("Upgrade-Insecure-Requests", "1")

	bow.Post("https://www.airfrance.fr/cgi-bin/AF/FR/fr/local/process/standardbooking/DisplayFlightPageAction.do",
		"application/x-www-form-urlencoded", strings.NewReader(string(jsonReq)))

	fmt.Println(bow.Body())

	bow.Dom().Find("tr.recommendation__overview").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})
}
