package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../core/util"

	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")

	util.CheckError(bow.Open("https://pay.flexve.com/botondepago/ravsa/inicio/"))

	fm, _ := bow.Form("[name='avail_search_form']")
	bow.Dom().Find("form").SetAttr("target", "_self")

	sessionID, _ := fm.Value("sesionId")
	milli := util.StrTimestampMilli()

	nextSessionID, _ := strconv.ParseInt(sessionID, 10, 32)
	nextSessionID = nextSessionID + 1

	fmt.Println(nextSessionID)

	fm.Set("departure_airport", "CCS")
	fm.Set("arrival_airport", "MAR")
	fm.Set("departure_date", "27-04-2018")
	fm.Set("arrival_date", "29-04-2018")
	fm.Set("passengers_ADT", "1")
	fm.Set("passengers_CNN", "0")
	fm.Set("passengers_INF", "0")
	fm.Set("action", "findFlights")
	fm.Set("new", "true")

	util.CheckError(fm.Submit())

	bow.Open("https://pay.flexve.com/botondepago/ravsa/formato_xml.php?sesionId=" + sessionID + "&sesionIP=201.209.100.100&departure_airport=CCS" +
		"&departure_date=27-04-2018&arrival_airport=MAR&passengers_ADT=1&passengers_CNN=0&passengers_INF=0&flightCode=&JsonData=&server=kiu" +
		"&airlineCode=AW&target=Production&_=" + milli)

	body := strings.Replace(util.Format(bow.Body()), ",\"itineraryOption\":null", "", -1)

	jqOut := util.ToJSONQ(body)

	itineraryOutArray, _ := jqOut.Array("travelInfo", "itinerary", "0", "itineraryOption")

	milli = util.StrTimestampMilli()

	bow.Open("https://pay.flexve.com/botondepago/ravsa/formato_xml.php?sesionId=" + string(nextSessionID) + "&sesionIP=201.209.100.100&departure_airport=MAR" +
		"&departure_date=29-04-2018&arrival_airport=CCS&passengers_ADT=1&passengers_CNN=0&passengers_INF=0&flightCode=&JsonData=&server=kiu" +
		"&airlineCode=AW&target=Production&_=" + milli)

	jqIn := util.ToJSONQ(bow.Body())
	itineraryIn, _ := jqIn.Object("travelInfo", "itinerary", "0", "itineraryOption", "0")
	itineraryOutArray[1] = itineraryIn

}
