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

	err := bow.Open("https://www.voeazul.com.br/uy/home")
	if err != nil {
		panic(err)
	}

	bow.AddRequestHeader("X-OneAgent-JS-Injection", "true")
	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")

	fm, _ := bow.Form("form.form-0001")
	_authkey, _ := fm.Value("_authkey_")

	fm.Set("orginalIata1", "SAO")
	fm.Set("destinationIata1", "SDU")
	fm.Set("departure1", "30/03/2018")
	fm.Set("arrival", "15/04/2018")
	fm.Set("ControlGroupSearch$SearchMainSearchView$RadioButtonMarketStructure", "RoundTrip")
	fm.Set("ControlGroupSearch$SearchMainSearchView$DropDownListPassengerType_ADT", "1")
	fm.Set("ControlGroupSearch$SearchMainSearchView$DropDownListPassengerType_CHD", "0")
	fm.Set("ControlGroupSearch$SearchMainSearchView$DropDownListPassengerType_INFANT", "0")
	fm.Set("origin1", "São Paulo - Todos os Aeroportos (SAO)")
	fm.Set("destination1", "Rio de Janeiro - Santos Dumont (SDU)")
	fm.Set("ControlGroupSearch$SearchMainSearchView$DropDownListSearchBy", "columnView")
	fm.Set("culture", "es-UY")
	fm.Set("ControlGroupSearch$SearchMainSearchView$TextBoxPromoCode", "CALLCENT")
	fm.Set("ControlGroupSearch$SearchMainSearchView$TextBoxMarketOrigin1", "São Paulo - Todos os Aeroportos (SAO)")
	fm.Set("ControlGroupSearch$SearchMainSearchView$CheckBoxUseMacOrigin1", "on")
	fm.Set("hdfSearchCodeDeparture1", "1N")
	fm.Set("ControlGroupSearch$SearchMainSearchView$TextBoxMarketDestination1", "Rio de Janeiro - Santos Dumont (SDU)")
	fm.Set("hdfSearchCodeArrival1", "1N")
	fm.Set("ControlGroupSearch$SearchMainSearchView$DropDownListMarketDay1", "30")
	fm.Set("ControlGroupSearch$SearchMainSearchView$DropDownListMarketMonth1", "2018-03")
	fm.Set("ControlGroupSearch$SearchMainSearchView$DropDownListMarketDay2", "15")
	fm.Set("ControlGroupSearch$SearchMainSearchView$DropDownListMarketMonth2", "2018-04")
	fm.Set("ControlGroupSearch$SearchMainSearchView$DropDownListFareTypes", "R")
	fm.Set("authkey_", _authkey)
	fm.Set("__EVENTTARGET", "ControlGroupSearch$LinkButtonSubmit")

	checkError(fm.Submit())

	bow.Dom().Find("tr.flight-info").Each(func(_ int, s *goquery.Selection) {
		fmt.Print(s.Html())
	})

}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
