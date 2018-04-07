package main

import (
	"fmt"
	"strings"

	"../../core/util"

	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()

	body := "No data found"
	for body == "No data found" {
		departure(bow)
		body = bow.Body()
	}

	body = "{\"results\":" + bow.Body() + "}"

	jq := util.ToJSONQ(body)

	exampleElement, _ := jq.Array("results")
	fmt.Println("Availability", exampleElement)

}

func departure(bow *browser.Browser) string {

	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")

	bow.AddRequestHeader("Content-Type", "text/html; charset=utf-8")
	util.CheckError(bow.Open("https://www.allegiantair.com/booking/AUS/CLE/2018-04-26/2018-05-18/1/null/nul"))

	var jsonReq = []byte(`{"outward":"2018-04-26","returning":"2018-05-18","destination_code":"CLE","origin_code":"AUS",
		"travelers":{"adult":1,"child_dobs":[],"lapchild_dobs":[]}}`)

	bow.Post("https://www.allegiantair.com/g4search/api/booking/search", "application/json", strings.NewReader(string(jsonReq)))

	manifestID := bow.State().Response.Header.Get("manifest-id")

	bow.AddRequestHeader("Accept", "application/json, text/javascript, */*; q=0.01")

	bow.Open("https://www.allegiantair.com/data/21/" + manifestID + "?id=" + manifestID)

	bow.AddRequestHeader("Accept", "application/json, text/javascript, */*; q=0.01")
	bow.AddRequestHeader("Manifest_Id", manifestID)
	bow.AddRequestHeader("G4-Client-Id", "booking-www")
	bow.AddRequestHeader("Host", "www.allegiantair.com")
	bow.AddRequestHeader("Referer", "https://www.allegiantair.com/booking/AUS/CLE/2018-04-26/2018-05-18/1/null/null")
	bow.AddRequestHeader("Silo-Id", "21")
	bow.AddRequestHeader("AJSUI-Version", "210.1.7")
	bow.AddRequestHeader("Connection", "keep-alive")
	bow.AddRequestHeader("X-Requested-With", "XMLHttpRequest")

	bow.Open("https://www.allegiantair.com/data/21/" + manifestID + "-" + "AUSCLE")

	return manifestID

}
