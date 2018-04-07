package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"../../core/util"

	"github.com/jmoiron/jsonq"
	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")

	util.CheckError(bow.Open("https://es.flyfrontier.com/sdbooking/Flight/InternalSelect?o1=ATL&d1=ORD&dd1=Apr%2020,%202018&dd2=Apr%2027,%202018" +
		"&ADT=1&r=true&umnr=false&mon=true&promo="))

	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	bow.AddRequestHeader("Accept_Language", "es-US,es-419;q=0.9,es;q=0.8")
	bow.AddRequestHeader("Connection", "keep-alive")
	bow.AddRequestHeader("Host", "es.flyfrontier.com")
	bow.AddRequestHeader("Upgrade-Insecure-Requests", "1")
	bow.AddRequestHeader("Referer", "https://es.flyfrontier.com/")

	bow.Open("https://es.flyfrontier.com/sdbooking/Flight/Select")

	jsonConent := getFlightData(bow.Body())

	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(jsonConent))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	exampleElement, _ := jq.Array("journeys")
	fmt.Println("FareData", exampleElement)
	//	fmt.Println(jsonConent)
}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}

func getFlightData(body string) string {
	var rgx = regexp.MustCompile(`FlightData = \'(.*?)\'`)
	rs := rgx.FindStringSubmatch(body)
	// Create replacer with pairs as arguments.
	r := strings.NewReplacer("FlightData = '", "", "}'", "}", "&quot;", "\"", "&amp;", "&")
	return r.Replace(rs[0])
}
