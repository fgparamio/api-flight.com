package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/jsonq"
	// "github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()

	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")

	checkError(bow.Open("https://tarmexico.com"))

	var jsonReq = []byte(`{"origin_airport":"ACA","destination_airport":"MTY","currency_origin":"MXN",
		"currency_destination":"MXN","transactionCurrency":"MXN","countryCode":"OM","departure_date":"2018-04-29",
		"return_date":"2018-05-10","adults":"1","children":"0","seniors":"0","infants":"0","promotional_code":""}`)

	milli := strconv.FormatInt(makeTimestampMilli(), 10)
	bow.Post("https://tarmexico.com/1.0/booking?i="+milli, "application/json; charset=UTF-8", strings.NewReader(string(jsonReq)))

	// TODO Move To Core
	data := map[string]interface{}{}
	body := strings.Replace(bow.Body(), "&#34;", "\"", -1)
	dec := json.NewDecoder(strings.NewReader(body))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	idBooking, _ := jq.String("content", "id_booking")

	fmt.Println(idBooking)

	milli = strconv.FormatInt(makeTimestampMilli(), 10)
	bow.Open("https://tarmexico.com/1.0/booking/" + idBooking + "/flights?i=" + milli + "&departure=2018-04-27&return=2018-04-30")

	// TODO Move To Core
	body = strings.Replace(bow.Body(), "&#34;", "\"", -1)
	dec = json.NewDecoder(strings.NewReader(body))
	dec.Decode(&data)
	jq = jsonq.NewQuery(data)

	flights, _ := jq.ArrayOfObjects("content", "flights", "departure_flights")
	fmt.Println("Flights:", flights)
}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}

func unixMilli(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func makeTimestampMilli() int64 {
	return unixMilli(time.Now())
}
