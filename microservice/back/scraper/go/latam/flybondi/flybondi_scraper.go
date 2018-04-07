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

	err := bow.Open("https://booking.flybondi.com/es/Flight/Search?adults=1&children=0&currency=ARS&departureDateString=2018-04-17" +
		"&fromCityCode=EPA&infants=0&returnDateString=2018-04-28&roundTrip=true&toCityCode=COR&useFlexDates=true")

	// Check GET Request
	util.CheckError(err)

	bodyHTML := bow.Body()
	securityToken := getSecurityToken(bodyHTML)

	var jsonReq = []byte(`{"interline":false,"fromCityCode":"EPA","toCityCode":"COR","departureDateString":"2018-04-17",
		"returnDateString":"2018-04-28","startDateStringOutbound":"2018-04-17","endDateStringOutbound":"2018-04-17",
		"startDateStringInbound":"2018-04-28","endDateStringInbound":"2018-04-28","adults":1,"children":0,"infants":0,
		"roundTrip":true,"useFlexDates":true,"isOutbound":true,"filterMethod":"100","promocode":"","currency":"ARS",
		"languageCode":"es-AR","fareTypeCategory":1,"IATANumber":"","securityToken":"`)

	securityTokenArray := []byte(securityToken + "\"}")
	request := append(jsonReq, securityTokenArray...)

	bow.Post("https://booking.flybondi.com/Api/AvailablityRequest/Post", "application/json; charset=UTF-8", strings.NewReader(string(request)))

	jq := util.ToJSONQ(bow.Body())

	exampleElement, _ := jq.Object("Availability", "OutboundSegments", "0", "LowestFareSummary")
	fmt.Println("First Segment", exampleElement)
}

func getSecurityToken(body string) string {
	position := strings.Index(body, "SecurityToken")
	token := body[position+15 : position+61]
	fmt.Println(token)
	return strings.Replace(token, "\"", "", -1)
}

func parseBodyWithInnerHTML(body string) string {
	body = strings.Replace(body, "charset=\"\\\"utf-8\\\"\"", "", -1)
	body = strings.Replace(body, "\"\\\"", "", -1)
	body = strings.Replace(body, "\\\"", "", -1)
	body = strings.Replace(body, "\">", "", -1)
	return body
}
