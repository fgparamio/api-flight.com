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

	util.CheckError(bow.Open("https://www.aerotucan.com/php/choosefly.php?tipo_vuelo=RT&origen=Huatulco&fecha_ida=2018-03-23" +
		"&destino=Oaxaca&fecha_vuelta=2018-03-31&adulto=1&menor=0&bebe=0&submit=Buscar+Vuelos+-%3E"))

	bow.Dom().Find("table.aerotucan2-article").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})
}
