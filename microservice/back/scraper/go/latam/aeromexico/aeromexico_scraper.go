package main

import (
	"fmt"
	"regexp"
	"strings"

	"../../core/util"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")

	util.CheckError(bow.Open("https://world.aeromexico.com/es/co/"))

	// Outputs: "reddit: the front page of the internet"
	fm, _ := bow.Form("form#flight")
	fm.Set("ob", "BOG")
	fm.Set("ib", "CUN")
	fm.Set("sd", "2018-05-13")
	fm.Set("ed", "2018-05-18")
	fm.Set("no", "Bogotá (BOG)")
	fm.Set("ni", "Cancún (CUN)")
	fm.Set("ti", "round")
	fm.Set("ad1", "1")
	fm.Set("ch1", "0")
	af, _ := fm.Value("af")

	util.CheckError(fm.Submit())
	modelData := getModelData(bow.Body())
	fmt.Println(modelData)

	bow.Post("http://co.aeromexico.com/SparksVista/FlightListOutboundQuote?ln=esp&cu=PO&sd=2018-05-13&ed=2018-05-18&ib=CUN"+
		"&ob=BOG&ni=Canc%C3%BAn%20(CUN)&no=Bogot%C3%A1%20(BOG)&alf=0&af="+af+"&ch1=0&ad1=1&ti=round&ac1=&iif=&cpf=&dsh=&ash=",
		"application/x-www-form-urlencoded; charset=UTF-8", strings.NewReader("model="+modelData))

	bow.Dom().Find("p.kdo-product--main-price").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}

func getModelData(body string) string {
	readLine := strings.Replace(body, "\n", "", -1)
	var rgx = regexp.MustCompile(` ModelData : \{(.*?)\};`)
	rs := rgx.FindStringSubmatch(readLine)
	r := strings.NewReplacer("};", "", " ModelData :", "")
	return r.Replace(rs[0])
}
