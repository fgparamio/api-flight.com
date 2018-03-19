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

	err := bow.Open("https://www.satena.com/")
	if err != nil {
		panic(err)
	}

	fm, _ := bow.Form("form#aym_frm_booking")
	fm.Set("boo_opt", "2")
	fm.Set("boo_ini", "BOG")
	fm.Set("boo_fin", "CLO")
	fm.Set("boo_dat_ini", "2018-03-31")
	fm.Set("boo_dat_fin", "2018-04-06")
	fm.Set("boo_adu", "1")
	fm.Set("boo_chi", "0")
	fm.Set("boo_inf", "0")
	fm.Set("boo_file_dat", "1")
	fm.Set("action", "S")

	checkError(fm.Submit())

	bow.Dom().Find("table.aym_table_booking").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
