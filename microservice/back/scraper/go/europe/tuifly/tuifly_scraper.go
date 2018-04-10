package main

import (
	"fmt"
	"regexp"
	"strings"

	"../../core/util"

	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()

	util.CheckError(bow.Open("https://en.tuifly.com/"))

	bow.Open("https://en.tuifly.com/flightoffers?origin=TXL&destination=MAD&start=2018-05-02&end=2018-05-26&adults=1")

	jq := util.ToJSONQ(getModelData(bow.Body()))

	outbounds, _ := jq.Object("outboundJourneys")
	fmt.Println("OutboundJourneys", outbounds)

	inbounds, _ := jq.Object("inboundJourneys")
	fmt.Println("InboundJourneys", inbounds)

}

func getModelData(body string) string {
	readLine := strings.Replace(body, "\n", "", -1)
	var rgx = regexp.MustCompile(`<script type=\"application/json\" class=\"js-trip-tile-list-data\">(.*?)</script>`)
	rs := rgx.FindStringSubmatch(readLine)
	r := strings.NewReplacer("<script type=\"application/json\" class=\"js-trip-tile-list-data\"> [", "", "]</script>", "")
	return r.Replace(rs[0])
}
