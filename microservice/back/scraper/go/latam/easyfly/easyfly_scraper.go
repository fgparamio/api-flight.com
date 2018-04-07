package main

import (
	"fmt"
	"time"

	"../../core/util"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetTimeout(2 * time.Minute)
	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	bow.AddRequestHeader("Connection", "keep-alive")
	bow.AddRequestHeader("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")
	bow.AddRequestHeader("Accept-Language", "es-US,es-419;q=0.9,es;q=0.8")
	bow.AddRequestHeader("Host", "www.easyfly.com.co")
	bow.AddRequestHeader("Referer", "https://www.easyfly.com.co/webcheckin/oldcheckin")
	bow.AddRequestHeader("Upgrade-Insecure-Requests", "1")

	util.CheckError(bow.Open("https://www.easyfly.com.co"))

	bow.Open("https://www.easyfly.com.co/flights?origins=1&originsText=Bogot%C3%A1+%28BOG%29&multi=&destinations=27" +
		"&originsTextReturn=Pereira+%28PEI%29&multiReturn=&flightType=1&departureDateEngine=31-03-2018" +
		"&returnDateEngine=29-04-2018&adt=1&chd=0&inf=0&promotionID=&tstPost=tstPost")

	// Parse Example
	bow.Dom().Find("span.txtPrice").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}
