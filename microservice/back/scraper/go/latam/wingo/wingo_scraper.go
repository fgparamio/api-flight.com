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

	err := bow.Open("https://www.wingo.com/es/Flight/Search?interline=false" +
		"&fromCityCode=BAQ&toCityCode=ADZ&departureDateString=2018-4-01" +
		"&returnDateString=2018-4-26&adults=1&children=0&infants=0" +
		"&roundTrip=true&useFlexDates=false&allInclusive=&promocode=" +
		"&fareTypes=&currency=COP")

	util.CheckError(err)

	bodyHTML := bow.Body()
	securityToken := getSecurityToken(bodyHTML)

	var jsonReq = []byte(`{"interline":false,"fromCityCode":"BAQ",
		"toCityCode":"ADZ","departureDateString":"2018-03-25","returnDateString":"2018-04-26",
		"startDateStringOutbound":"2018-03-25","endDateStringOutbound":"2018-03-25",
		"startDateStringInbound":"2018-04-26","endDateStringInbound":"2018-04-26",
		"adults":1,"children":0,"infants":0,"roundTrip":true,"useFlexDates":true,"isOutbound":true,
		"filterMethod":"100","promocode":"","currency":"COP","languageCode":"es-CO","fareTypeCategory":1,
		"IATANumber":"","securityToken":"`)

	securityTokenArray := []byte(securityToken + "\"}")
	request := append(jsonReq, securityTokenArray...)

	bow.Post("https://www.wingo.com/Api/AvailablityRequest/Post", "application/json; charset=UTF-8", strings.NewReader(string(request)))

	jq := util.ToJSONQ(bow.Body())

	exampleElement, _ := jq.Object("Availability")
	fmt.Println("Availability", exampleElement)

}

func getSecurityToken(body string) string {

	position := strings.Index(body, "SecurityToken")
	token := body[position+15 : position+61]
	fmt.Println(token)
	return strings.Replace(token, "\"", "", -1)
}
