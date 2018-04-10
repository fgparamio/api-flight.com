package main

import (
	"fmt"
	"strings"

	"../../core/util"
	"github.com/PuerkitoBio/goquery"

	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")
	util.CheckError(bow.Open("https://www.flypgs.com/en"))

	fm, _ := bow.Form("form#fligth-searh")
	fm.Set("DEPPORT", "AMS")
	fm.Set("ARRPORT", "ESB")
	fm.Set("LBLDEPPORT", "Amsterdam")
	fm.Set("LBLARRPORT", "Ankara")
	fm.Set("TRIPTYPE", "27/04/2018")
	fm.Set("DEPDATE", "0")
	fm.Set("RETDATE", "12/05/2018")
	fm.Set("ADULT", "1")
	fm.Set("CHILD", "0")
	fm.Set("INFANT", "0")
	fm.Set("STUDENT", "0")
	fm.Set("SOLDIER", "0")
	fm.Set("CURRENCY", "GBP")
	fm.Set("LC", "EN")
	fm.Set("FLEX", "")
	fm.Set("resetErrors", "T")
	fm.Set("clickedButton", "btnSearch")
	fm.Set("DEPDATEO", "17")
	fm.Set("RETDATEO", "32")

	util.CheckError(fm.Submit())

	bow.Post("https://book.flypgs.com/Common/MemberRezvResults.jsp?activeLanguage=EN", "application/x-www-form-urlencoded",
		strings.NewReader("DEPPORT=AMS&ARRPORT=ESB&LBLDEPPORT=Amsterdam&LBLARRPORT=Ankara&TRIPTYPE=R&DEPDATE=27%2F04%2F2018"+
			"&RETDATE=30%2F04%2F2018&ADULT=1&CHILD=0&INFANT=0&STUDENT=0&SOLDIER=0&CURRENCY=GBP&LC=EN&FLEX=&"+
			"resetErrors=T&clickedButton=btnSearch&DEPDATEO=17&RETDATEO=20"))

	bow.Dom().Find("span.flightPrice").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}
