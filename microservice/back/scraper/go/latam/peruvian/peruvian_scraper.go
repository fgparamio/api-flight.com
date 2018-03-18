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
	err := bow.Open("http://www.peruvian.pe/pe")
	if err != nil {
		panic(err)
	}

	// Outputs: "reddit: the front page of the internet"
	fm, _ := bow.Form("form#form_cn")
	fm.Set("origen", "LIM")
	fm.Set("destino", "CUZ")
	fm.Set("tipoViaje", "R")
	fm.Set("fec_salida", "31.03-2018")
	fm.Set("fec_retorno", "28-04-2018")
	fm.Set("num_adulto", "1")
	fm.Set("num_menor", "0")
	fm.Set("num_infante", "0")

	checkError(fm.Submit())

	// Parse Example
	bow.Dom().Find("table.table-wizard").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
