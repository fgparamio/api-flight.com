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

	fm := bow.Forms()[0] 
	fm.Set("name", "0")
	fm.Set("data_input_origin", "MEX")
	fm.Set("data_input_destination", "CUN")
	fm.Set("hf_countrycode_origin", "MX")
	fm.Set("hf_countrycode_destination", "MX")
	fm.Set("data_CBX_origin", "0")
	fm.Set("data_CBX_destination", "0")
	fm.Set("data_HotelTab_destination", "CUM")
	fm.Set("PromoCode", "")
	fm.Set("tripType", "0")
	fm.Set("input_data_from", "03/22/2018")
	fm.Set("input_data_to", "03/24/2018")
	fm.Set("txAdultValue", "1")
	fm.Set("txMinorValue", "0")
	fm.Set("txInfantValue", "0")
	fm.Set("txPassportValue", "1")
	fm.Set("dniValue", "0")



	bow.Dom().Find("div#js_availability_container").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
