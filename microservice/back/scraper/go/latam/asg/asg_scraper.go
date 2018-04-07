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
	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	bow.AddRequestHeader("Host", "www.asg.com.mx")
	util,CheckError(bow.Open("https://www.asg.com.mx"))
	bow.Dom().Find("form").AddClass("classForm")

	fm, _ := bow.Form("form.classForm")
	fm.Set("input-type", "r")
	fm.Set("input-origen", "ciudad obregon (CEN)")
	fm.Set("input-destino", "Guerrero Negro (GRN)")
	fm.Set("input-fIda", "22/03/2018")
	fm.Set("input-fVuelta", "24/03/2018")
	fm.Set("input-paxs", "1")
	fm.Set("input-menor", "0")
	fm.Set("input-infante", "0") 

	util.CheckError(fm.Submit())

	bow.Dom().Find("div.margen-top-mini").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}
