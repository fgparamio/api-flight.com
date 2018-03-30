package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")

	checkError(bow.Open("https://www.spirit.com/Default.aspx"))

	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xm…plication/xml;q=0.9,*/*;q=0.8")
	bow.AddRequestHeader("Accept_Language", "es-US,es-419;q=0.9,es;q=0.8")
	bow.AddRequestHeader("Connection", "keep-alive")
	bow.AddRequestHeader("Referer", "https://www.easyjet.com/es/")
	bow.AddRequestHeader("Host", "www.spirit.com")
	bow.AddRequestHeader("Upgrade-Insecure-Requests", "1")
	bow.AddRequestHeader("Origin", "https://www.spirit.com")
	bow.AddRequestHeader("Content-Type", "application/x-www-form-urlencoded")
	bow.AddRequestHeader("Referer", "https://www.spirit.com/Default.aspx")

	bow.Post("https://www.spirit.com/Default.aspx?action=search", "application/x-www-form-urlencoded",
		strings.NewReader("bypassHC=False&birthdates=&lapoption=&awardFSNumber=&bookingType=F&hotelOnlyInput=&autoCompleteValueHidden"+
			"=&carPickUpTime=16&carDropOffTime=16&tripType=roundTrip&vacationPackageType=on&from=ATL&to=BOS&departDate=05%2F03%2F2018"+
			"&departDateDisplay=05%2F03%2F2018&returnDate=05%2F06%2F2018&returnDateDisplay=05%2F06%2F2018&ADT=1&CHD=0&INF="+
			"0&promoCode=&fromMultiCity1=&toMultiCity1=&dateMultiCity1=&dateMultiCityDisplay1=&fromMultiCity2=&toMultiCity2=&dateMultiCity2"+
			"=&dateMultiCityDisplay2=&fromMultiCity3=&toMultiCity3=&dateMultiCity3"+
			"=&dateMultiCityDisplay3=&fromMultiCity4=&toMultiCity4=&dateMultiCity4=&dateMultiCityDisplay4=&redeemMiles=false"))

	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xm…plication/xml;q=0.9,*/*;q=0.8")
	bow.AddRequestHeader("Host", "www.spirit.com")
	bow.AddRequestHeader("Accept-Language", "es-US,es-419;q=0.9,es;q=0.8")
	bow.AddRequestHeader("Referer", "https://www.spirit.com/Default.aspx")
	bow.AddRequestHeader("Cache-Control", "max-age=0")
	bow.AddRequestHeader("Connection", "keep-alive")
	bow.Open("https://www.spirit.com/DPPCalendarMarket.aspx")

	bow.Dom().Find("div.cal-date").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
