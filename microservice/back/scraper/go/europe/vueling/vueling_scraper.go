package main

import (
	"fmt"

	"github.com/headzoo/surf"
)

func main() {

	// create a new collector
	// c := colly.NewCollector(
	// 	colly.Debugger(&debug.LogDebugger{}),
	// 	colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36"),
	// )

	// c.SetRequestTimeout(90 * time.Second)

	// // attach callbacks after login
	// c.OnRequest(func(r *colly.Request) {

	// 	// buf := new(bytes.Buffer)
	// 	// buf.ReadFrom(r.Body)
	// 	// b := buf.Bytes()
	// 	// s := *(*string)(unsafe.Pointer(&b))

	// 	// println(s)
	// 	// println(len(s))

	// 	r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	// 	r.Headers.Set("Accept-Language", "es-US,es-419;q=0.9,es;q=0.8")
	// 	r.Headers.Set("Host", "tickets.vueling.com")
	// 	r.Headers.Set("Connection", "keep-alive")
	// 	r.Headers.Set("Cache-Control", "max-age=0")
	// 	r.Headers.Set("Content-Type", "application/x-www-form-urlencoded")
	// 	r.Headers.Set("Content-Length", "1353")
	// 	r.Headers.Set("Host", "tickets.vueling.com")
	// 	r.Headers.Set("Origin", "https://tickets.vueling.com")
	// 	r.Headers.Set("Referer", "https://tickets.vueling.com/")
	// 	r.Headers.Set("Upgrade-Insecure-Requests", "1")
	// 	r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")
	// 	r.Headers.Set("Cookie", "ASP.NET_SessionId=zomntwb2gdlgbb0uraj25lk2; DC=n01; s_gts=1; _yldr_ab=a; skysales=!HGOgG7pLpApMpSacbeTUc6RteVAc63eEEI7GD8fJrAgIDDb8fcfHrc/0i/eyRIll3Y9K3UOKiMBJgSs=; s_fid=1DBFA18BDFCA61D9-072150E33C20CD6D; ttc_sec=1522714966568; idioma=es-ES; rxVisitor=1522714996595N3F6MV7NLLMNV3O9IGQRKTPMDR7E3SBA; _ga=GA1.2.717543385.1522715000; _gid=GA1.2.154530404.1522715000; abTesting=id=93&value=HGm8DvJ56WcHpuFSiQaqB%2fXS6EgY6%2fD0%2f9JfHI%2ffBIw%3d; s_cc=true; newhome=false; s_vi=[CS]v1|2D61679A0507AEC9-400001112000C84E[CE]; __qca=P0-1999485746-1522716459515; OriginStation=Station=BCN&Country=CO; personalizacion=id=93&value=N%2fl4fgU7KW6dX5REMIJOdW6t0adj46jcVpAxX9jc5edzhYtyfahWGw%3d%3d; CookiePolicy=IsSession=False&AcceptDateTime=636583149703083750; Currency=Currency=EUR&Country=CO&IsAutomatic=False; ttc_flag=true; cto_lwid=48685195-a612-4518-9b0f-42dab02d0c4b; s_nr=1522718764869; ak_bmsc=FFE5F8E187A20339360D5C97F472DD7817D46C43E11700009EACC35ACF108406~pladLrLx6gCPiB5qQ3Sg11u/Dh1XXcP3fNUIO460mjLTEExRSEgBfx6QWMDLUztjhVsePjHnN+47WOfDcq4OmOGNnbi2m5ltG3bnZzv0yRJnboUHbG6wv0FGTpV6fwTQnOaYd6ei8fv7Vem6/AOa8sQSRyb/DHDlOTXqoJ/Hgif91mcEmWOn6VCS/r0a1+zxaZ1mQdhEkWXWAN8QbbHmY44eMe5fDJztckb1PXZzHjimZqutGFlL4VdhstzNsUGViD; _yldr_history=www.vueling.com%7CNO_REFERRER%7Ctickets.vueling.com; _yldr_session_nr=4; busquedaexpress=id=93&value=nocSZk7LrL%2bUTN6ZGDuFw2VQbza5ApCkGJCA7COvlEKHPWk%2bCXW4P5XDIRJ0ze6P9KQGSNmqVXC1n0cjZig%2bsB06HR69VchW7MPN6d7tiP8%3d; ttc=1522775933523; _yldr_user_fq=32; dtLatC=2; dtPC=3$575931309_356h-vCRLMPJKHAIIHIDGBOEBBFEDDFHIDGENBGC; dtCookie=3$AF8FE04AE3F233D3662026119646B615|RUM+Default+Application|1; dtSa=true%7CC%7C-1%7CVolver%20al%20inicio%7C-%7C1522778971366%7C575931309_356%7Chttps%3A%2F%2Ftickets.vueling.com%2FScheduleSelect.aspx%7CVueling%7C1522775940734%7C; rxvt=1522780772266|1522777741550; mbox=check#true#1522779034|session#1522778973861-972307#1522780834; _gat_b47894630a27ca274f2b7dca99ccfa7f=1; _uetsid=_uet8a347ebe; bm_sv=23F9EB10BB93A8682132BAF025CB13B4~gb0qDuqfWyeWD7a4g5Q6bRAfEkAtQznkx6icmCOOTRvvhMPMsFalT84EpUbWgsQaQ0dGrduHvNGij+Hw3MXTv6S34ECHw0nAiMQyIuUOaAYqThhAvqYgD3xdDJjh5L0RvDzcHanIrGBJxEVSJtTRGyeuxsuENW6ZRTtpUnmLHB4=; s_cpmh=%5B%5B%27TraficoDirecto%27%2C%271522778980815%27%5D%5D; s_sq=vuelingprod%3D%2526c.%2526a.%2526activitymap.%2526page%253DPagina%252520de%252520Inicio%2526link%253DBuscar%252520vuelos%2526region%253Dtab-flights%2526pageIDType%253D1%2526.activitymap%2526.a%2526.c%2526pid%253DPagina%252520de%252520Inicio%2526pidt%253D1%2526oid%253Dfunctiononclick%252528event%252529%25257BreturnSKYSALES.Util.validate%252528this%25252Cnull%25252Cnull%25252C%252527no_summary%252527%252529%25253B%25257D%2526oidt%253D2%2526ot%253DA")
	// })

	// err := c.Post("https://tickets.vueling.com/XmlSearch.aspx",
	// 	map[string]string{
	// 		"__EVENTTARGET":   "AvailabilitySearchInputXmlSearchView$LinkButtonNewSearch",
	// 		"__EVENTARGUMENT": "",
	// 		"__VIEWSTATE":     "",
	// 		"pageToken":       "",
	// 		"AvailabilitySearchInputXmlSearchView$RadioButtonMarketStructure": "RoundTrip",
	// 		"AvailabilitySearchInputXmlSearchView$TextBoxMarketOrigin1":       "Barcelona",
	// 		"AvailabilitySearchInputXmlSearchView$TextBoxMarketDestination1":  "Londres (Gatwick)",
	// 		"date_picker": "2018-04-21",
	// 		"AvailabilitySearchInputXmlSearchView$DropDownListMarketDay1":           "21",
	// 		"AvailabilitySearchInputXmlSearchView$DropDownListMarketMonth1":         "2018-04",
	// 		"AvailabilitySearchInputXmlSearchView$DropDownListMarketDay2":           "29",
	// 		"AvailabilitySearchInputXmlSearchView$DropDownListMarketMonth2":         "2018-04",
	// 		"AvailabilitySearchInputXmlSearchView$DropDownListPassengerType_ADT":    "1",
	// 		"AvailabilitySearchInputXmlSearchView$DropDownListPassengerType_CHD":    "0",
	// 		"AvailabilitySearchInputXmlSearchView$DropDownListPassengerType_INFANT": "0",
	// 		"AvailabilitySearchInputXmlSearchView$ResidentFamNumSelector":           "",
	// 		"AvailabilitySearchInputXmlSearchView$ExtraSeat":                        "",
	// 		"AvailabilitySearchInputXmlSearchView$DropDownListSearchBy":             "columnView",
	// 		"departureStationCode1":                                                 "BCN",
	// 		"arrivalStationCode1":                                                   "LGW",
	// 		"ErroneousWordOrigin1":                                                  "",
	// 		"SelectedSuggestionOrigin1":                                             "",
	// 		"ErroneousWordDestination1":                                             "",
	// 		"SelectedSuggestionDestination1":                                        "",
	// 		"departureStationCode2":                                                 "",
	// 		"arrivalStationCode2":                                                   "",
	// 		"ErroneousWordOrigin2":                                                  "",
	// 		"SelectedSuggestionOrigin2":                                             "",
	// 		"ErroneousWordDestination2":                                             "",
	// 		"SelectedSuggestionDestination2":                                        ""})

	// if err != nil {
	// 	// log.Fatal(err)
	// }

	// // attach callbacks after login
	// c.OnResponse(func(r *colly.Response) {
	// 	// fmt.Println("response received", string(r.Body))
	// })

	// c.Visit("https://tickets.vueling.com/")

	// start scraping
	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")

	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	bow.AddRequestHeader("Host", "www.vueling.com")
	bow.AddRequestHeader("Cache-Control", "max-age=0")
	bow.AddRequestHeader("Connection", "keep-alive")

	checkError(bow.Open("http://www.vueling.com/es"))

	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	bow.AddRequestHeader("Accept-Language", "es-US,es-419;q=0.9,es;q=0.8")
	bow.AddRequestHeader("Content-Type", "application/x-www-form-urlencoded")
	// bow.AddRequestHeader("Content-Length", "1353")
	bow.AddRequestHeader("Host", "tickets.vueling.com")
	bow.AddRequestHeader("Origin", "http://www.vueling.com")
	bow.AddRequestHeader("Referer", "http://www.vueling.com/es")
	bow.AddRequestHeader("Cache-Control", "max-age=0")
	bow.AddRequestHeader("Connection", "keep-alive")

	fm, _ := bow.Form("form#SkySales")

	fm.Set("__EVENTTARGET", "AvailabilitySearchInputXmlSearchView$LinkButtonNewSearch")
	fm.Set("__EVENTARGUMENT", "")
	fm.Set("__VIEWSTATE", "")
	fm.Set("pageToken", "")
	fm.Set("AvailabilitySearchInputXmlSearchView$RadioButtonMarketStructure", "RoundTrip")
	fm.Set("AvailabilitySearchInputXmlSearchView$TextBoxMarketOrigin1", "Londres")
	fm.Set("AvailabilitySearchInputXmlSearchView$TextBoxMarketDestination1", "Barcelona")
	fm.Set("date_picker", "2018-04-21")
	fm.Set("AvailabilitySearchInputXmlSearchView$DropDownListMarketDay1", "21")
	fm.Set("AvailabilitySearchInputXmlSearchView$DropDownListMarketMonth1", "2018-04")
	fm.Set("AvailabilitySearchInputXmlSearchView$DropDownListMarketDay2", "29")
	fm.Set("AvailabilitySearchInputXmlSearchView$DropDownListMarketMonth2", "2018-04")
	fm.Set("AvailabilitySearchInputXmlSearchView$DropDownListPassengerType_ADT", "1")
	fm.Set("AvailabilitySearchInputXmlSearchView$DropDownListPassengerType_CHD", "0")
	fm.Set("AvailabilitySearchInputXmlSearchView$DropDownListPassengerType_INFANT", "0")
	fm.Set("AvailabilitySearchInputXmlSearchView$ResidentFamNumSelector", "")
	fm.Set("AvailabilitySearchInputXmlSearchView$ExtraSeat", "")
	fm.Set("AvailabilitySearchInputXmlSearchView$DropDownListSearchBy", "columnView")
	fm.Set("departureStationCode1", "LON")
	fm.Set("arrivalStationCode1", "BCN")
	fm.Set("ErroneousWordOrigin1", "")
	fm.Set("SelectedSuggestionOrigin1", "")
	fm.Set("ErroneousWordDestination1", "")
	fm.Set("SelectedSuggestionDestination1", "")
	fm.Set("departureStationCode2", "")
	fm.Set("arrivalStationCode2", "")
	fm.Set("ErroneousWordOrigin2", "")
	fm.Set("SelectedSuggestionOrigin2", "")
	fm.Set("ErroneousWordDestination2", "")
	fm.Set("SelectedSuggestionDestination2", "")

	checkError(fm.Submit())

	cookies := ""
	for _, cookie := range bow.State().Request.Cookies() {
		cookies = cookies + cookie.Name + "=" + cookie.Value + "; "
	}

	fmt.Println(cookies)

	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	bow.AddRequestHeader("Accept-Language", "es-US,es-419;q=0.9,es;q=0.8")
	bow.AddRequestHeader("Cache-Control", "max-age=0")
	bow.AddRequestHeader("Connection", "keep-alive")
	bow.AddRequestHeader("Host", "tickets.vueling.com")
	bow.AddRequestHeader("Referer", "http://www.vueling.com/es")
	bow.AddRequestHeader("Upgrade-Insecure-Requests", "1")

	// TODO change this line
	bow.AddRequestHeader("Cookie", "ASP.NET_SessionId=zomntwb2gdlgbb0uraj25lk2; DC=n01; s_gts=1; _yldr_ab=a; skysales=!HGOgG7pLpApMpSacbeTUc6RteVAc63eEEI7GD8fJrAgIDDb8fcfHrc/0i/eyRIll3Y9K3UOKiMBJgSs=; s_fid=1DBFA18BDFCA61D9-072150E33C20CD6D; ttc_sec=1522714966568; idioma=es-ES; rxVisitor=1522714996595N3F6MV7NLLMNV3O9IGQRKTPMDR7E3SBA; _ga=GA1.2.717543385.1522715000; _gid=GA1.2.154530404.1522715000; abTesting=id=93&value=HGm8DvJ56WcHpuFSiQaqB%2fXS6EgY6%2fD0%2f9JfHI%2ffBIw%3d; s_cc=true; newhome=false; s_vi=[CS]v1|2D61679A0507AEC9-400001112000C84E[CE]; __qca=P0-1999485746-1522716459515; OriginStation=Station=BCN&Country=CO; personalizacion=id=93&value=N%2fl4fgU7KW6dX5REMIJOdW6t0adj46jcVpAxX9jc5edzhYtyfahWGw%3d%3d; CookiePolicy=IsSession=False&AcceptDateTime=636583149703083750; Currency=Currency=EUR&Country=CO&IsAutomatic=False; ttc_flag=true; cto_lwid=48685195-a612-4518-9b0f-42dab02d0c4b; s_nr=1522718764869; ak_bmsc=FFE5F8E187A20339360D5C97F472DD7817D46C43E11700009EACC35ACF108406~pladLrLx6gCPiB5qQ3Sg11u/Dh1XXcP3fNUIO460mjLTEExRSEgBfx6QWMDLUztjhVsePjHnN+47WOfDcq4OmOGNnbi2m5ltG3bnZzv0yRJnboUHbG6wv0FGTpV6fwTQnOaYd6ei8fv7Vem6/AOa8sQSRyb/DHDlOTXqoJ/Hgif91mcEmWOn6VCS/r0a1+zxaZ1mQdhEkWXWAN8QbbHmY44eMe5fDJztckb1PXZzHjimZqutGFlL4VdhstzNsUGViD; _yldr_history=www.vueling.com%7CNO_REFERRER%7Ctickets.vueling.com; _yldr_session_nr=4; busquedaexpress=id=93&value=nocSZk7LrL%2bUTN6ZGDuFw7J9rPLZ8hz2lEPQdczqvKYCN1qG62a8sUWdkswf6M0C0%2fwlEHVwXBZhebPguEAdz0q7zhKtv%2bwf; ttc=1522778987047; _yldr_user_fq=34; dtPC=3$578985126_338h-vCRLMPJKHAIIHIDGBOEBBFEDDFHIDGENBGC; dtLatC=170; dtCookie=3$AF8FE04AE3F233D3662026119646B615|RUM+Default+Application|1; dtSa=true%7CC%7C-1%7CVolver%20al%20inicio%7C-%7C1522779726535%7C578985126_338%7Chttps%3A%2F%2Ftickets.vueling.com%2FScheduleSelect.aspx%7CVueling%7C1522778993921%7C; rxvt=1522781527271|1522777741550; mbox=session#1522778973861-972307#1522781588|check#true#1522779788; _gat_b47894630a27ca274f2b7dca99ccfa7f=1; _uetsid=_uet8a347ebe; bm_sv=23F9EB10BB93A8682132BAF025CB13B4~gb0qDuqfWyeWD7a4g5Q6bRAfEkAtQznkx6icmCOOTRvvhMPMsFalT84EpUbWgsQaQ0dGrduHvNGij+Hw3MXTv6S34ECHw0nAiMQyIuUOaAbtBTcA8sFyY5n4wkm/dWJX7Qo12NUaj/KEXRsB3ZnGdB7k2yoFYwHikS/jZHFdRvk=; s_cpmh=%5B%5B%27TraficoDirecto%27%2C%271522779732446%27%5D%5D; s_sq=vuelingprod%3D%2526c.%2526a.%2526activitymap.%2526page%253DPagina%252520de%252520Inicio%2526link%253DBuscar%252520vuelos%2526region%253Dtab-flights%2526pageIDType%253D1%2526.activitymap%2526.a%2526.c%2526pid%253DPagina%252520de%252520Inicio%2526pidt%253D1%2526oid%253Dfunctiononclick%252528event%252529%25257BreturnSKYSALES.Util.validate%252528this%25252Cnull%25252Cnull%25252C%252527no_summary%252527%252529%25253B%25257D%2526oidt%253D2%2526ot%253DA")

	checkError(bow.Open("https://tickets.vueling.com/ScheduleSelect.aspx"))
	fmt.Println(bow.Body())
}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
