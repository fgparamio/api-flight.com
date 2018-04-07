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

	util.CheckError(bow.Open("https://book.andesonline.com"))

	util.CheckError(bow.Open("https://book.andesonline.com/FlightSearch?o1=AEP&d1=TUC&dd1=2018-03-29&dd2=2018-04-25&r=true&ADT=1&CHD=0&inl=0"))
	util.CheckError(bow.Open("https://book.andesonline.com/Flight/InternalSelect?o1=AEP&d1=TUC&dd1=2018-03-29&dd2=2018-04-25&r=true&ADT=1&CHD=0&inl=0&s=true&mon=true"))

	bow.Open("https://book.andesonline.com/Flight/Select")

	bow.Dom().Find("div#js_availability_container").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}
