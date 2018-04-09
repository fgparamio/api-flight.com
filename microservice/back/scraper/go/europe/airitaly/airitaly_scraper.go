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
	util.CheckError(bow.Open("https://www.airitaly.com/it-IT/travel-plan/book/booking-search-result?adt=1&chd=0&inf=0&dep=FCO" +
		"&dep_dt=20180519&rt=MXP&rt_dt=20180531&ow=false&cls=Y&ct=false"))

	bow.AddRequestHeader("Accept-Encoding", "deflate, br")

	bow.Post("https://www.airitaly.com/api/booking/GetAirPrice", "application/x-www-form-urlencoded; charset=UTF-8",
		strings.NewReader("adt=1&chd=0&inf=0&dep=FCO&dep_dt=20180519&rt=MXP&rt_dt=20180531&ow=false&cls=Y&ct=false"))

	fmt.Println(bow.Body())
}
