package main

import (
	"fmt"

	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()

	checkError(bow.Open("http://www.viajaconsol.com/ventas/index.php"))

	bow.Open("http://www.viajaconsol.com/ventas/init.php?preset=1&origen_cod=4&destino_cod=3&fecha_ida=13/04/2018" +
		"&fecha_vuelta=21/04/2018&adultos=1&ninos=0&bebes=0&incluye_vuelta=1&ahora_string=2018-03-22T11:06:54&width=932" +
		"&height=974&info_referer=http://www.viajaconsol.com/index.php")

	bow.Open("http://www.viajaconsol.com/ventas/index.php?pagina=paso2")
	bow.Open("http://www.viajaconsol.com/ventas/index.php")
	bow.Open("http://www.viajaconsol.com/ventas/ajax/control_refresh_cuerpo.php?pagina=paso2")
	bow.Open("http://www.viajaconsol.com/ventas/ajax/refresh_cuerpo.php?pagina=paso2")

	fmt.Println(bow.Dom().Html())

}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
