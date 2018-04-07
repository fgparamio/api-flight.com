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

	util.CheckError(bow.Open("https://albatrosair.aero"))

	bow.Open("https://albatrosair.aero/reservas/busqueda.php?location_origin=CCS&location_destination=SJO" +
		"&time_depart=2018-05-09&time_return=2018-05-16&pax_CNN=0&pax_ADT=1&pax_INF=0&is_roundtrip=1")

	bow.Open("https://albatrosair.aero/reservas/ajaxbuscar2.php?acc=buscar&is_roundtrip=1&location_origin=CCS" +
		"&location_destination=SJO&time_depart=2018-05-09&time_return=2018-05-16&pax_ADT=1&pax_CNN=0&pax_INF=0&flexible=&cabin_type=")

	// Parse Example
	bow.Dom().Find("table#tbsalida").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

	bow.Dom().Find("table#tbretorno").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}
