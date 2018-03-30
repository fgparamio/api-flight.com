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

	bow.Open("https://www.easyjet.com/deeplink?lang=ES&dep=MAD&dest=LIS&dd=2018-05-02&rd=2018-05-13&apax=1&cpax=0&ipax=0&SearchFrom=SearchPod2_/es/&dt=Desktop&isOneWay=off&pid=www.easyjet.com")
	bow.Open("https://www.easyjet.com/es/buy/flights?isOneWay=off&pid=www.easyjet.com")

	// TODO: Try delete this line. 
	bow.AddRequestHeader("Cookie", "_ga=GA1.3.1509590120.1521332889; lang2012=es-es; ejCC_3=v=4941825604653417075&i=-8586802740263873429; CMSPOD=dub-sc1-blue; cookies.js=1; idb*=%7B%22departure%22%3Anull%2C%22when%22%3Anull%2C%22flag%22%3Anull%2C%22price%22%3Anull%2C%22destination%22%3Anull%2C%22fromDate%22%3Anull%2C%22toDate%22%3Anull%2C%22customDate%22%3Anull%2C%22mapLocation%22%3Anull%2C%22moveMap%22%3Anull%2C%22region%22%3A%22es-es%22%7D; RBKPOD=dub-rbk-blue; AVLPOD=dub-avl-blue; labi=036cd0da-4c74-faaa-1431-2cfd322ff4ea; odb*=MAD; FunnelQuery=%7b%22OriginIata%22%3a%22MAD%22%2c%22DestinationIata%22%3a%22LIS%22%2c%22OutboundDate%22%3a%222018-03-31%22%2c%22ReturnDate%22%3a%222018-04-13%22%2c%22OutboundFlightNumber%22%3anull%2c%22ReturnFlightNumber%22%3anull%2c%22FunnelJourney%22%3a%22default%22%2c%22NumberOfAdults%22%3a1%2c%22NumberOfChildren%22%3a0%2c%22NumberOfInfants%22%3a0%2c%22OpenSearchPanel%22%3afalse%2c%22ShowFlexiFares%22%3afalse%2c%22RemainOnStep1%22%3afalse%2c%22ComponentSender%22%3a%22SearchPod2_%2fes%2f%22%2c%22CurrencyCode%22%3anull%2c%22PaymentTypeCode%22%3anull%7d; ej20SearchCookie=ej20Search_0=MAD|LIS|2018-03-31T00:00:00|2018-04-13T00:00:00|1|0|0|False||0|2018-03-29 12:25:46Z; ej20RecentSearches=ej20RecentSearch_0=MAD|LIS|2018-03-31T00:00:00|2018-04-13T00:00:00|1|0|0|False||0|2018-03-29 12:25:46Z; FunnelJourney=%7B%22CorrelationToken%22%3A%22C02B1347-92D9-AD58-CBC4-71AA94B26371%22%2C%22PaymentType%22%3Anull%2C%22Origin%22%3A%22MAD%22%2C%22Destination%22%3A%22LIS%22%2C%22Outbound%22%3A%222018-03-31%22%2C%22Return%22%3A%222018-04-13%22%2C%22Component%22%3A%22SearchPod2_%2Fes%2F%22%2C%22SearchPanelOpen%22%3Afalse%2C%22FlexiFares%22%3Afalse%2C%22Adults%22%3A1%2C%22Children%22%3A0%2C%22Infants%22%3A0%2C%22CurrencyCode%22%3A%22EUR%22%2C%22CurrencyPriority%22%3A2%2C%22JourneyPairId%22%3A1%2C%22CarSearchQuery%22%3Anull%2C%22Pairs%22%3A%5B%5D%2C%22FunnelJourney%22%3A%22default%22%7D; _abck=86960B32DFBEA516E8DAC3F480BA1943B8339755ED7A0000A3B2AD5A37F30B79~0~joih8/G8cbMsSRWrV+FUxilPOWKjHYp4dZXgtZSFP0I=~-1~-1; bm_sz=98187877EF7ABD8368A376935A596857~QAAQBjcZyJ7FSm5iAQAA45hzd1QR23Gbzefs5VTEKoLHRd2pvnsDWpKowhPsCGVyNzJaSd6gbnKj0bappRYA55+qYvQFBsXS4PXscIwFkeaohdlU/ApgqtDXpv03L77QWubBl8XDRqzRAFTjozC4tXYRmnxTjBA+AML+A+/OL3p+MVQirz1YmNS5jhsOIw==; _gid=GA1.3.564282885.1522422482; mmapi.store.p.0=%7B%22mmparams.d%22%3A%7B%7D%2C%22mmparams.p%22%3A%7B%22pd%22%3A%221553958482768%7C%5C%22-930928818%7CIwAAAApVAgDMUH29%2Fg8AAREAAUKLP67dBgB6jlQMUJbVSNHI0idnjNVIAAAAAP%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FAAZEaXJlY3QBCBAFAAEAAAAAAAAARbIBAEWyAQBFsgEAAAAAAAAAAAFF%5C%22%22%2C%22srv%22%3A%221553958482771%7C%5C%22lvsvwcgeu01%5C%22%22%2C%22uat%22%3A%221553958482844%7C%7B%5C%22Domain%5C%22%3A%5C%22es%5C%22%2C%5C%22DepAirport%5C%22%3A%5C%22mad%5C%22%2C%5C%22Destination%5C%22%3A%5C%22lis%5C%22%2C%5C%22TripDuration%5C%22%3A%5C%228to14%5C%22%2C%5C%22LeadTime%5C%22%3A%5C%220to3%5C%22%2C%5C%22PassengerQty%5C%22%3A%5C%221%5C%22%2C%5C%22PassengerType%5C%22%3A%5C%22single%5C%22%2C%5C%22SaturdayNight%5C%22%3A%5C%22true%5C%22%7D%22%7D%7D; mmapi.store.s.0=%7B%22mmparams.d%22%3A%7B%7D%2C%22mmparams.p%22%3A%7B%7D%7D; bm_mi=DA1CBCABCCEA4F70D93F9FA411C8B9FD~ViBnMdjklZzYR9OjMiTzYmtA8Cz6lTwmWOn2TwBsDaYmzlrZCHWlWkCOcBiceIul5AS6/LiD/hyXoUE8yec2ofPILKc9amhVqR2ZSBmU48lM59yH4JQiHJhDCU/OhvEL3c6ByqiCTPSG0rHe72P/yjg4LPU7s9/igVq3pLJOqEN/FYyKzkJbgpCc6cvN2QMNGYaOWZ5yCcrVt0gDme6BTpMoPWzAn2KOBpakmL8kTRyAdxnEhscyOM33fWV7ECDeLlkDVR9gLt9OHPq4ajoKcw==; ak_bmsc=1F63FD663173370D772DE47FC8379742C8193706FE4E0000D752BE5A36F43D6F~pl2rrNyqxOxc1ueSTF3crTcmkSJ1skMr2ZbPR/MIRwJsQ6KQgTHie2ilDL2jpORt22dsf8syoKq6zQ6y6k0AUS5v4ZDuzVeVSJ2WywaYlzawyHqfekbP1ge6MQuX7Ll/7POhK8iDw+L+gYLNQviPK1qsbFbxMXtp2E6QHn/Aied2x2RpvkRRSpWap8yTh6brZ8zXDWr64kpH0mDSTuhhFt6Y+f9FKlb7BqOHEf8GiDMQ1FtkZYo6P6aeFbDa7eYQad; 47620=; bm_sv=0AEA741F4BEB5E70BDD5BDA614371797~Ioi7FD9iFIgODp/t5GLz/HkEr/eb4GL7uN+r6LvUg68yAsIJEEAX564X6Mpuri1SFpQ7Mrhc4MDNQNWPUMZkutNLApaDZ6YWMJdA3JajBZxoNOXcvsR47LeG403+pV/vZVq6wg3UpMFFZi9C4C/hr37LYCOk/YiODkoYwqDk8Rc=; akacd_TrueClarity_SC=1523027346~rv=80~id=b8174e0190be4f73de0a9552e9de98c7; ADRUM=s=1522422545146&r=https%3A%2F%2Fwww.easyjet.com%2Fes%2F%3F0")

	checkError(bow.Open("https://www.easyjet.com/ejavailability/api/v16/availability/query?AdditionalSeats=0&AdultSeats=1&ArrivalIata=LIS&ChildSeats=0" +
		"&DepartureIata=MAD&IncludeAdminFees=true&IncludeFlexiFares=false&IncludeLowestFareSeats=true&IncludePrices=true&Infants=0&IsTransfer=false" +
		"&LanguageCode=ES&MaxDepartureDate=2018-05-01&MaxReturnDate=2018-06-14&MinDepartureDate=2018-04-30&MinReturnDate=2018-06-12"))

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
