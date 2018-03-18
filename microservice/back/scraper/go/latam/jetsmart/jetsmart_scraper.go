package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")
	err := bow.Open("https://booking.jetsmart.com/Flight/InternalSelect?" +
		"o1=SCL&d1=LIM&dd1=2018-3-25&ADT=1&CHD=0&inl=0&r=true&s=true&mon=true&cur=CLP&pc=&dd2=2018-3-30")
	// Check GET Request
	checkError(err)
	// Parse Example
	bow.Dom().Find("div#js_availability_container").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})
}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
