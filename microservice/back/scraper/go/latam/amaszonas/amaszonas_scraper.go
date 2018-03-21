package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")

	checkError(bow.Open("https://www.amaszonas.com/es-bo"))

	bow.AddRequestHeader("Content-Type", "application/x-www-form-urlencoded")

	bow.Dom().Find("form#reservas").SetAttr("target", "_self")
	fm, _ := bow.Form("[name='reservas']")
	fm.Set("origen", "VVI")
	fm.Set("destino", "ASU")
	fm.Set("fecIda", "30/03/2018")
	fm.Set("horaIda", "")
	fm.Set("fechaida", "")
	fm.Set("fecVuelta", "11/04/2018")
	fm.Set("horaVuelta", "")
	fm.Set("fechaVuelta", "")
	fm.Set("adultos", "1")
	fm.Set("menores", "0")
	fm.Set("bebes", "0")
	fm.Set("cabina", "Y")
	fm.Set("tipoviaje", "2")

	checkError(fm.Submit())

	bow.AddRequestHeader("Host", "reservas2.amaszonas.com")
	bow.AddRequestHeader("Referer", "https://www.amaszonas.com/es-bo/")

	getQuery := "https://reservas2.amaszonas.com/reservasiwtg/avail.php?i=1&tipo=2&origen_v=VVI&destino_v=ASU" +
		"&fechaidaf=30/03/2018&fechaida=&fechaVueltaf=11/04/2018&fechaVuelta=&adultos=1&menores=0&bebes=0&pais=BO"

	bow.Open(getQuery)

	fm2, _ := bow.Form("[name='reserva']")
	bow.Dom().Find("form").SetAttr("target", "_self")
	it, _ := fm2.Value("it")

	data := "it=" + it + "&tipo_new=1&fechaaux=30/03/2018&actualiza=1"

	bow.Post("https://reservas2.amaszonas.com/reservasiwtg/server/avail_ajax.php?cc="+getCC(),
		"application/x-www-form-urlencoded", strings.NewReader(data))

	bow.Dom().Find("li.columna_origendestino").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

	data = "it=" + it + "&tipo_new=2&fechaaux=11/04/2018"

	bow.Post("https://reservas2.amaszonas.com/reservasiwtg/server/avail_ajax.php?cc="+getCC(),
		"application/x-www-form-urlencoded", strings.NewReader(data))

	bow.Dom().Find("li.columna_origendestino").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})
}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}

func getCC() string {
	now := time.Now()
	return strconv.Itoa(now.Hour()) + "7" + strconv.Itoa(now.Minute()) + "7" + strconv.Itoa(now.Second()) +
		"7" + strconv.Itoa(now.Nanosecond() * 1000)[0:1]
}
