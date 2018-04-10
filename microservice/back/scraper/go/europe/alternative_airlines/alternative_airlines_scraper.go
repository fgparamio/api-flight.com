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
	util.CheckError(bow.Open("https://www.alternativeairlines.com/dba"))

	util.CheckError(bow.Open("https://www.alternativeairlines.com/flights/MAD-PAR?outbound=2018-04-29&return=2018-05-11" +
		"&adults=1&children=0&infants=0&class=Y&currency=COP"))

	stToken := getSt(bow.Body())

	bow.Post("https://www.alternativeairlines.com/api/search?api=1", "application/x-www-form-urlencoded; charset=UTF-8",
		strings.NewReader("st="+stToken))

	jq := util.ToJSONQ(bow.Body())

	exampleElement, _ := jq.Object("priceGroups")
	fmt.Println("Availability", exampleElement)
}

func getSt(body string) string {
	position := strings.Index(body, "st:")
	token := body[position+5 : position+37]
	fmt.Println(token)
	return strings.Replace(token, "\"", "", -1)
}
