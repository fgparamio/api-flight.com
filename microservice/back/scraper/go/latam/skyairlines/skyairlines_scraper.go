package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jmoiron/jsonq"
	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")

	err := bow.Open("https://www.skyairline.com/otros/flujo-compra/busqueda-vuelos?fromCityCode=EZE&fromCityCode=SCL"+
		"&toCityCode=SCL&toCityCode=EZE&departureDateString=2018-03-24&departureDateString=2018-04-27&fareTypeCategory=1"+
		"&adults=1&children=0&infants=0&pets=0&currency=USD&isMobileSearch=false&isNewSearch=true")

	if err != nil {
		panic(err)
	}

	getSecurityToken(bow.Body())

	var jsonReq = []byte(`{"interline":false,"fromCityCode":["EZE","SCL"],
		"toCityCode":["SCL","EZE"],"departureDate":"2018-03-24","departureDateString":["2018-03-24","2018-04-27"],
		"returnDateString":"2018-03-27","startDateStringOutbound":["2018-03-21","2018-04-24"],
		"endDateStringOutbound":["2018-03-27","2018-04-30"],"startDateStringInbound":"2018-03-21",
		"endDateStringInbound":"2018-03-27","adults":1,"children":0,"infants":0,"pets":0,"roundTrip":true,
		"useFlexDates":true,"isOutbound":true,"filterMethod":"101","promocode":"","currency":"USD",
		"useRealCacheKey":false,"regularSearch":false,"taLogin":"","taPin":"","iataNumber":"","apiKey":"",
		"fareTypeCategory":1,"languageCode":"es-ES","securityToken":"","isMobileSearch":false}`)

	bow.Post("https://www.skyairline.com/Availability/Post", "application/json; charset=UTF-8", strings.NewReader(string(jsonReq)))

	data := map[string]interface{}{}
	body := strings.Replace(bow.Body(), "&#34;", "\"", -1)
	dec := json.NewDecoder(strings.NewReader(body))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	exampleElement, _ := jq.Object("Availability")
	fmt.Println("Availability", exampleElement)

}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}

func getSecurityToken(body string) string {
	position := strings.Index(body, "window.securityToken")
	token := body[position : position+25]
	fmt.Println(token)
	return strings.Replace(token, "\"", "", -1)
}
