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

	err := bow.Open("http://compre2.voegol.com.br/search.aspx?culture=pt-br")
	if err != nil {
		panic(err)
	}

	// Outputs: "reddit: the front page of the internet"
	fm, _ := bow.Form("form#SkySales")
	fm.Set("__EVENTTARGET", "")
	fm.Set("__EVENTARGUMENT", "")
	fm.Set("__VIEWSTATE", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$DropDownListResidentCountry", "BR")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$DropDownListSearchBy", "columnView")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawDepartureStation_0", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawArrivalStation_0", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawDepartureDate_0", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawDepartureStation_1", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawArrivalStation_1", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawDepartureDate_1:", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawDepartureStation_2", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawArrivalStation_2", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawArrivalStation_2", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawDepartureStation_3", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawArrivalStation_3", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawDepartureDate_3", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawDepartureStation_4", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawArrivalStation_4", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawDepartureDate_4", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawDepartureStation_5", "")	
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawArrivalStation_5", "")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOpenJawDepartureDate_5", "")

	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$RadioButtonMarketStructure", "RoundTrip")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketOrigin1", "Rio de Janeiro - Todos os aeroportos (RIO)")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$TextBoxMarketDestination1", "São Paulo - Todos os aeroportos (SAO)")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$DropDownListMarketDay1", "03")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$DropDownListMarketMonth1", "2018-04")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$DropDownListMarketDay2", "28")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$DropDownListMarketMonth2", "2018-04")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$DropDownListPassengerType_ADT", "1")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$DropDownListPassengerType_CHD", "0")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$DropDownListPassengerType_INFT", "0")
	fm.Set("ControlGroupSearchView$AvailabilitySearchInputSearchView$textBoxPromoCode", "")
	fm.Set("pnrSearchType", "loc")

	checkError(fm.Submit())
	
	// Parse Example
	bow.Dom().Find("span.fareValue").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
