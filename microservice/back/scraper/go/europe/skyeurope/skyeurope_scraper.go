package main

import (
	"fmt"
	"strings"

	"../../core/util"

	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")
	util.CheckError(bow.Open("http://www.skyeurope.com"))

	bow.Dom().Find("form").AddClass("classForm")
	fm, _ := bow.Form("form.classForm")
	fm.Set("task", "search")
	fm.Set("lang", "en")
	fm.Set("start", "MAD")
	fm.Set("arrive", "BCN")
	fm.Set("people", "1")
	fm.Set("front_day", "09")
	fm.Set("front_month", "04")
	fm.Set("front_year", "2018")
	fm.Set("back_day", "16")
	fm.Set("back_month", "04")
	fm.Set("sa", "0")
	fm.Set("debug", "0")

	util.CheckError(fm.Submit())

	bow.Post("http://www.skyeurope.com/functions/ajax.php", "application/x-www-form-urlencoded; charset=UTF-8",
		strings.NewReader("task=search&lang=en&start=MAD&arrive=BCN&people=1&front_day=09"+
			"&front_month=04&front_year=2018&back_day=16&back_month=04&back_year=2018&sa=0&debug=0"))

	fmt.Println(bow.Body())
}
