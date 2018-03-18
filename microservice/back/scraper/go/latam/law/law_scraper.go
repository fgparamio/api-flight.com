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
	err := bow.Open("https://www.vuelalaw.com/cl")
	if err != nil {
		panic(err)
	}

	// Outputs: "reddit: the front page of the internet"
	fm, _ := bow.Form("form#booking")
	fm.Set("tipoviaje", "0")
	fm.Set("codigoorigen", "SCL")
	fm.Set("codigodestino", "PAP")
	fm.Set("actualciudad", "")
	fm.Set("lang", "SP")
	fm.Set("desde", "SCL")
	fm.Set("hasta", "PAP")
	fm.Set("from", "28/03/2018")
	fm.Set("to", "04/04/2018")
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
