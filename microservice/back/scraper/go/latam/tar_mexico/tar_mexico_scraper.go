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

	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")

	util.CheckError(bow.Open("https://tarmexico.com"))

	var jsonReq = []byte(`{"origin_airport":"ACA","destination_airport":"MTY","currency_origin":"MXN",
		"currency_destination":"MXN","transactionCurrency":"MXN","countryCode":"OM","departure_date":"2018-04-29",
		"return_date":"2018-05-10","adults":"1","children":"0","seniors":"0","infants":"0","promotional_code":""}`)

	milli := util.StrTimestampMilli()
	bow.Post("https://tarmexico.com/1.0/booking?i="+milli, "application/json; charset=UTF-8", strings.NewReader(string(jsonReq)))

	jq := util.ToJSONQ(bow.Body())

	idBooking, _ := jq.String("content", "id_booking")

	fmt.Println(idBooking)

	milli = util.StrTimestampMilli()
	bow.Open("https://tarmexico.com/1.0/booking/" + idBooking + "/flights?i=" + milli + "&departure=2018-04-27&return=2018-04-30")

	jq = util.ToJSONQ(bow.Body())

	flights, _ := jq.ArrayOfObjects("content", "flights", "departure_flights")
	fmt.Println("Flights:", flights)
}
