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
	util.CheckError(bow.Open("https://www.eurowings.com/es.html"))

	fm, _ := bow.Form("form.form")
	fm.Set("originName", "Madrid")
	fm.Set("origin", "MAD")
	fm.Set("airlineData", "")
	fm.Set("destinationName", "Hamburgo")
	fm.Set("destination", "HAM")
	fm.Set("outboundDateTextInput", "29.04.18")
	fm.Set("inboundDateTextInput", "16.05.18")
	fm.Set("maxDate", "29.04.18 - 16.05.18")
	fm.Set("dateDate", "2018-04-29")
	fm.Set("toDate", "2018-05-16")
	fm.Set("roundTrip", "true")
	fm.Set("adults", "1")
	fm.Set("childs", "0")
	fm.Set("infants", "0")
	fm.Set("tpid", "")
	fm.Set("language", "es-ES")

	util.CheckError(fm.Submit())

	bow.Open("https://www.eurowings.com/skysales/Deeplink.aspx?culture=es-ES&o=MAD&d=BRU&od=2018-04-22&adt=1&chd=0&inf=0&lng=es-ES&rd=2018-05-09&t=r")

	bow.Open("https://www.eurowings.com/skysales/GlobalSelect.aspx?culture=es-ES")

	bow.Dom().Find("div.input-radio-a11y").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}
