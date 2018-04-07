package main

import (
	"fmt"

	"../../core/util"

	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()

	util.CheckError(bow.Open("https://kibe-om.kiusys.net"))
	bow.Dom().Find("form").AddClass("classForm")
	fm, _ := bow.Form("form.classForm")
	token, _ := fm.Value("csrfmiddlewaretoken")

	fm.Set("origin_autocomplete", "")
	fm.Set("origin", "Q0JC")
	fm.Set("destination_autocomplete", "LA PAZ (LPB), EL ALTO AIRPORT (LPB), Bolivia")
	fm.Set("destination", "TFBC")
	fm.Set("departure_date", "27/3/2018")
	fm.Set("return_date", "31/3/2018")
	fm.Set("adults", "1")
	fm.Set("minors", "0")
	fm.Set("infants", "0")

	util.CheckError(fm.Submit())

	bow.Open("https://kibe-om.kiusys.net/availability/")
	bow.Open("https://kibe-om.kiusys.net/api/week/?arrival_date=27-3-2018&return_date=31-3-2018")

	fmt.Println(token)

	bow.AddRequestHeader("referer", "https://kibe-om.kiusys.net/availability/")
	bow.AddRequestHeader("x-csrftoken", token)
	bow.Open("https://kibe-om.kiusys.net/api/get_availability/?adults=1&departure_date=27-3-2018&destination=LPB" +
		"&direct=false&infants=0&minors=0&origin=CBB&return_date=31-3-2018")

	jq := util.ToJSONQ(bow.Body())

	exampleElement, _ := jq.Array("going")
	fmt.Println("Going", exampleElement)
	exampleElement, _ = jq.Array("return")
	fmt.Println("Return", exampleElement)

}
