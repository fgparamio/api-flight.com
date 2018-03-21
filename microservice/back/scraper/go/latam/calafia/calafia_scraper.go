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

	checkError(bow.Open("https://www.calafiaairlines.com"))

	bow.Dom().Find("input#complete_origen_value").SetAttr("value", "Cabo San Lucas")
	bow.Dom().Find("input#complete_destino_value").SetAttr("value", "La Paz")
	bow.Dom().Find("input#date_from").SetAttr("value", "03/04/2018")
	bow.Dom().Find("input#date_to").SetAttr("value", "20/04/2018")
	bow.Dom().Find("span.lbl-defaulto").SetText("1 Adulto")

	bow.Click("a.btn-search")

	bow.Open("https://www.calafiaairlines.com/Reservaciones/")	

	bow.Dom().Find("div.lista-vuelos").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})
}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
