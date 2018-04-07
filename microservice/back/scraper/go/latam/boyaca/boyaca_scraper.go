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

	util.CheckError(bow.Open("http://www.aeroboyaca.com/reservas"))

	fm, _ := bow.Form("[name='formulario']")
	fm.Set("tipotrayecto", "Ida y Regreso")
	fm.Set("origen", "Paipa")
	fm.Set("destino", "Bogotá")
	fm.Set("fecha", "2018-03-29")
	fm.Set("fecha2", "2018-04-05")
	fm.Set("cantpasajeros", "1")
	fm.Set("menores2", "0")
	fm.Set("codigopromocion", "")
	fm.Set("Submit", "Buscar")

	util.CheckError(fm.Submit())

	bow.Dom().Find("table.vuelos").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}
