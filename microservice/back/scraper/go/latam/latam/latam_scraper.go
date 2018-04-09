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

	err := bow.Open("http://booking.lan.com/es_co/apps/personas/compra?fecha1_dia=31&fecha1_anomes=2018-04" +
		"&fecha2_dia=14&fecha2_anomes=2018-05&from_city2=SCL&to_city2=BOG&auAvailability=1&ida_vuelta=ida_vuelta&vuelos_origen=Bogot%C3%A1" +
		"&from_city1=BOG&vuelos_destino=Santiago%20de%20Chile&to_city1=SCL&flex=1" +
		"&vuelos_fecha_salida_ddmmaaaa=31/04/2018&vuelos_fecha_regreso_ddmmaaaa=14/05/2018&cabina=Y&nadults=1&nchildren=0&ninfants=0")

	// Check GET Request
	util.CheckError(err)

	var jsonReq = []byte(`{"language":"ES","country":"CO","portal":"personas","application":"compra_normal","section":"step2",
		"cabin":"Y","adults":1,"children":0,"infants":0,"roundTrip":true,"departureDate":"2018-04-31","origin":"BOG",
		"destination":"SCL","referenceFlight":{"departureDate":"2018-05-14","destination":"BOG","origin":"SCL"},"tripType":"outbound"}`)

	bow.Post("http://booking.lan.com/ws/booking/quoting/fares_availability/5.0/rest/get_availability", "application/json; charset=UTF-8",
		strings.NewReader(string(jsonReq)))

	jq := util.ToJSONQ(bow.Body())

	exampleElement, _ := jq.Object("data", "itinerary", "routesMap")
	fmt.Println("First Segment", exampleElement)
}
