package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func main() {

	// create a new collector
	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36"),
	)

	// attach callbacks after login
	c.OnRequest(func(r *colly.Request) {

		// r.Headers.Set(":authority", "booking.airasia.com")
		// r.Headers.Set(":method", "GET")
		// r.Headers.Set(":path", "/Flight/Select?o1=PEK&d1=CNX&culture=en-GB&dd1=2018-04-13&dd2=2018-04-20&r=true&ADT=1&s=true&mon=true&cc=USD&c=false")
		// r.Headers.Set(":scheme", "https")
		// r.Headers.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
		// r.Headers.Set("referer", "https://www.airasia.com/en/home.page")
		// r.Headers.Set("accept-language", "es-US,es-419;q=0.9,es;q=0.8")
		// r.Headers.Set("upgrade-insecure-requests", "1")
		// r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")
	})

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response received", string(r.Body))
		// c.Visit("https://booking.airasia.com/Flight/Select?o1=PEK&d1=CNX&culture=en-GB&dd1=2018-04-13&dd2=2018-04-20&r=true&ADT=1&s=true&mon=true&cc=USD&c=false")
	})

	c.Visit("https://www.airasia.com/en/home.page")
	c.Visit("https://booking.airasia.com/Flight/Select?o1=PEK&d1=CNX&culture=en-GB&dd1=2018-04-13&dd2=2018-04-20&r=true&ADT=1&s=true&mon=true&cc=USD&c=false")
}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
