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
	bow.Open("https://www.easyjet.com")

	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xm…plication/xml;q=0.9,*/*;q=0.8")
	bow.AddRequestHeader("Accept_Language", "es-US,es-419;q=0.9,es;q=0.8")
	bow.AddRequestHeader("Connection", "keep-alive")
	bow.AddRequestHeader("Referer", "https://www.easyjet.com/es/")
	bow.AddRequestHeader("Host", "www.easyjet.com")
	bow.AddRequestHeader("Upgrade-Insecure-Requests", "1")
	bow.AddRequestHeader("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")

	bow.Open("https://www.easyjet.com/deeplink?lang=ES&dep=MAD&dest=LIS&dd=2018-03-31&rd=2018-04-13&apax=1&cpax=0&ipax=0&SearchFrom=SearchPod2_/es/&dt=Desktop&isOneWay=off&pid=www.easyjet.com")
	bow.Open("https://www.easyjet.com/es/buy/flights?isOneWay=off&pid=www.easyjet.com")
	bow.Post("https://www.easyjet.com/ejrebooking/api/v16/funnel/emptybasket", "application/json;charset=utf-8", strings.NewReader(""))
	bow.Open("https://www.easyjet.com/ejrebooking/api/v16/account/getloggedinstatus")
	bow.Open("https://www.easyjet.com/ejcms/nocache/jsCallbacks/NearestAirport.aspx")

	var jsonReq = []byte(`{"Data":"036cd0da-4c74-faaa-1431-2cfd322ff4ea/MAD/LIS/1/0/0/False/ES/EUR/1"}`)

	bow.AddRequestHeader("Accept", "application/json, text/plain, */*")
	bow.AddRequestHeader("ADRUM", "isAjax:true")
	bow.Post("https://www.easyjet.com/ejavailability/api/v16/analytics/sendsearch", "application/json;charset=utf-8", strings.NewReader(string(jsonReq)))

	bow.AddRequestHeader("Accept", "application/json, text/plain, */*")
	bow.AddRequestHeader("Accept-Language", "es-US,es-419;q=0.9,es;q=0.8")
	bow.AddRequestHeader("Cookie", "_ga=GA1.3.1509590120.1521332889; lang2012=es-es; ejCC_3=v=4941825604653417075&i=-8586802740263873429;"+"CMSPOD=dub-sc1-blue; _gid=GA1.3.19802652.1522277252; cookies.js=1; idb*=%7B%22departure%22%3Anull%2C%22when%22%3Anull%2C%22flag%22%3Anull%2C%22price"+"%22%3Anull%2C%22destination%22%3Anull%2C%22fromDate%22%3Anull%2C%22toDate%22%3Anull%2C%22customDate%22%3Anull%2C%22mapLocation%22%3Anull%2C%22moveMap"+"%22%3Anull%2C%22region%22%3A%22es-es%22%7D; RBKPOD=dub-rbk-blue; WPPOD=2; AVLPOD=dub-avl-blue; labi=036cd0da-4c74-faaa-1431-2cfd322ff4ea; odb*=MAD;"+" bm_sz=0DC445C5192E1658107519F6DC21C2F7~QAAQK/"+"BPFxQHPkliAQAAf024cVZ7KV6o5r90u2JUCmrSrgytA7nVNZkmuNmnFgAdxbOccoKtTwU92l8ZStLpBJ3jcJNPxuDgWuX6SbcI0oTdQhv6uFBUlAt7zzg7JW0Ydvi0WGh8MVjl2dZ9d67lRBLNu3Lh8sN1PKcs6XoaBlHpDKDgsYSzugufBxCHFG3V;"+" bm_mi=74664BFB3468D632CD67B63892AE3DC6~ZhaVg9Uu/22vRu4PZpcpJjmimogPm7Q3JPdgPTS7iIJ2yFmZad5FyP1/h+CTq/4Mu5F6WLKABAYFiQGAa4lW5mCCHsL+JNQgHvwdeQDGbq2tsuM0ssUt/7vMANp7z8vBFHo/"+"J90bF7vmk0NA1tMYikF5SgH/EgcYPad9TARCzCA0N2X3ZPpm7aouRjD6cYIDkXbaBuicXhfFb/8heHGnBBl1ZWMfRR2X7VOODRjAYsBlYDaqsPC8lIig74eCHCqTsW8sqL/MeR209KYL8yh4AQ==;"+" ak_bmsc=9E0EAF403439DBA6BCB3F062C4B73EC6174FF02B522E000036DBBC5ABCF0836F~plJV+tZnGYjQMQUWAyFPhnDwdkUv9pMn6OuLemHW4CcMv4CCVnkof8GOAvOeIz+b/lGduTJNgPa+TBcmmXd72X9nAow1NL3Q/"+"Jy45h00CQMw/ONQZhrAd3iuGooCrnbKWB5Xko660D3/o6iblC60thRXNTRzJK000Te618ChjGEn3OBwR17E1xUchi64GL80OzNMUCAeHmaj61xAPlFiXtGtFNrFXQPxi/BHJJY95aZXt/whMFUb7G3GjB13gZ8XbH; 47620=;"+" ADRUM=s=1522326342771&r=https%3A%2F%2Fwww.easyjet.com%2Fes%2F%3F0; FunnelQuery=%7b%22OriginIata%22%3a%22MAD%22%2c%22DestinationIata%22%3a%22LIS%22%2c%22OutboundDate%22%3a%222018-03-31"+"%22%2c%22ReturnDate%22%3a%222018-04-13%22%2c%22OutboundFlightNumber%22%3anull%2c%22ReturnFlightNumber%22%3anull%2c%22FunnelJourney%22%3a%22default%22%2c%22NumberOfAdults%22%3a1%2c%22"+"NumberOfChildren%22%3a0%2c%22NumberOfInfants%22%3a0%2c%22OpenSearchPanel%22%3afalse%2c%22ShowFlexiFares%22%3afalse%2c%22RemainOnStep1%22%3afalse%2c%22ComponentSender%22%3a%22SearchPod2_%2"+"fes%2f%22%2c%22CurrencyCode%22%3anull%2c%22PaymentTypeCode%22%3anull%7d; ej20SearchCookie=ej20Search_0=MAD|LIS|2018-03-31T00:00:00|2018-04-13T00:00:00|1|0|0|False||0|2018-03-29 12:25:46Z;"+" ej20RecentSearches=ej20RecentSearch_0=MAD|LIS|2018-03-31T00:00:00|2018-04-13T00:00:00|1|0|0|False||0|2018-03-29 12:25:46Z; akacd_TrueClarity_SC=1522931152~rv=80~id=981cfbda183c9a207e59800d3da4085f;"+" FunnelJourney=%7B%22CorrelationToken%22%3A%22C02B1347-92D9-AD58-CBC4-71AA94B26371%22%2C%22PaymentType%22%3Anull%2C%22Origin%22%3A%22MAD%22%2C%22Destination%22%3A%22LIS%22%2C%22Outbound%22%3A%222018-03-31"+"%22%2C%22Return%22%3A%222018-04-13%22%2C%22Component%22%3A%22SearchPod2_%2Fes%2F%22%2C%22SearchPanelOpen%22%3Afalse%2C%22FlexiFares%22%3Afalse%2C%22Adults%22%3A1%2C%22Children%22%3A0%2C%22Infants%22%3A0%2C%22"+"CurrencyCode%22%3A%22EUR%22%2C%22CurrencyPriority%22%3A2%2C%22JourneyPairId%22%3A1%2C%22CarSearchQuery%22%3Anull%2C%22Pairs%22%3A%5B%5D%2C%22FunnelJourney%22%3A%22default%22%7D;"+" mmapi.store.p.0=%7B%22mmparams.d%22%3A%7B%7D%2C%22mmparams.p%22%3A%7B%22pd%22%3A%221553862347145%7C%5C%22-982901655%7CIQAAAApVAgDMUH29%2Fg8AAREAAUKmSxiABQCPsQ02cJXVSNHI0idnjNVIAAAAAP"+"%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FAAZEaXJlY3QBCBAEAAEAAAAAAAAARbIBAEWyAQBFsgEAAAAAAAAAAAFF%5C%22%22%2C%22srv%22%3A%221553862347149%7C%5C%22lvsvwcgeu04%5C%22%22%2C%22uat%22%3A%221553862348607"+"%7C%7B%5C%22Domain%5C%22%3A%5C%22es%5C%22%2C%5C%22DepAirport%5C%22%3A%5C%22mad%5C%22%2C%5C%22Destination%5C%22%3A%5C%22lis%5C%22%2C%5C%22TripDuration%5C%22%3A%5C%228to14%5C%22%2C%5C%22"+"LeadTime%5C%22%3A%5C%220to3%5C%22%2C%5C%22PassengerQty%5C%22%3A%5C%221%5C%22%2C%5C%22PassengerType%5C%22%3A%5C%22single%5C%22%2C%5C%22SaturdayNight%5C%22%3A%5C%22true%5C%22%7D%22%7D%7D;"+" mmapi.store.s.0=%7B%22mmparams.d%22%3A%7B%7D%2C%22mmparams.p%22%3A%7B%7D%7D; _abck=86960B32DFBEA516E8DAC3F480BA1943B8339755ED7A0000A3B2AD5A37F30B79~-1~m2HEYsUmkcdcig+aBiEOqpOCIl8xbexDHlLDdTmEabM=~0~-1;"+" bm_sv=D4A561C14A43F580C1F3A51A04FFD145~pa2axUne2bXYQ221iOjtLa0v2SCZZf/Y6oDuhXUREhgm9gH6OCQ9WPSQHmmltAKSrsxFsYzKcR9T+oXnq4fjyHrwTOSG5DozhQPqUb8Q1rZF5vTmzSBxjXQW9lIwy5b0gWD5eUE1v8v2Dt7u0Ci5ARw+b+xrK6z+LLtM/X0JLMk=")

	checkError(bow.Open("https://www.easyjet.com/ejavailability/api/v16/availability/query?AdditionalSeats=0&AdultSeats=1&ArrivalIata=LIS&ChildSeats=0" +
		"&DepartureIata=MAD&IncludeAdminFees=true&IncludeFlexiFares=false&IncludeLowestFareSeats=true&IncludePrices=true&Infants=0&IsTransfer=false" +
		"&LanguageCode=ES&MaxDepartureDate=2018-04-01&MaxReturnDate=2018-06-14&MinDepartureDate=2018-03-30&MinReturnDate=2018-06-12"))

	siteCookies := bow.SiteCookies()
	for _, cookie := range siteCookies {
		fmt.Println(cookie.Name)
	}

	data := map[string]interface{}{}
	body := strings.Replace(bow.Body(), "&#34;", "\"", -1)
	dec := json.NewDecoder(strings.NewReader(body))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	exampleElement, _ := jq.Array("AvailableFlights")
	fmt.Println("Availability", exampleElement)
}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
