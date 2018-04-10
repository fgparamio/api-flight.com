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
	util.CheckError(bow.Open("https://www.atlasglb.com/en"))

	fm, _ := bow.Form("[name='formTicket']")
	fm.Set("from_input", "DUBLIN, Dublin  Airport (DUB), IRELAND(REPUBLIC OF)")
	fm.Set("from", "DUB")
	fm.Set("to_input", "ISTANBUL, Ataturk Airport (IST), TURKEY")
	fm.Set("to", "IST")
	fm.Set("lang", "EN")
	fm.Set("direction", "0")
	fm.Set("depdate", "25/05/2018")
	fm.Set("retdate", "03/06/2018")
	fm.Set("adult", "1")
	fm.Set("yp", "0")
	fm.Set("chd", "0")
	fm.Set("inf", "0")
	fm.Set("sc", "0")
	fm.Set("stu", "0")
	fm.Set("tsk", "0")
	fm.Set("refid", "REFERERDELETE")
	fm.Set("paramstatus", "1")
	fm.Set("openjaw", "")
	fm.Set("bannerSize", "200x200")

	util.CheckError(fm.Submit())

	bow.Dom().Find("table.bookingTable").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}
