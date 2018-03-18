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

	checkError(bow.Open("http://www.starperu.com/es"))

	// Outputs: "reddit: the front page of the internet"
	fm, _ := bow.Form("form#frm_paso1")
	fm.Set("tipoviaje", "R")
	fm.Set("codigoorigen", "LIM")
	fm.Set("codigodestino", "CUZ")
	fm.Set("lang", "SP")
	fm.Set("desde", "LIM")
	fm.Set("hasta", "CUZ")
	fm.Set("from", "29/03/2018")
	fm.Set("to", "23/04/2018")
	fm.Set("tld3", "1 Adulto")
	fm.Set("tld1", "0 Niño")
	fm.Set("tld2", "0 Bebes")
	fm.Set("botoncillo2", "buscar")

	checkError(fm.Submit())

	// Parse Example
	bow.Dom().Find("table").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
