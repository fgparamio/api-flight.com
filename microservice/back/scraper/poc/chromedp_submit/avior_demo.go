// Command submit is a chromedp example demonstrating how to fill out and
// submit a form.
package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	// var res string
	err = c.Run(ctxt, scrap())
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	// log.Printf("got: `%s`", res)
}

func scrap() chromedp.Tasks {

	return chromedp.Tasks{
		chromedp.Navigate("https://www.aviorair.com"),

		// chromedp.SendKeys(`#cityOrigin`, kb.Enter, chromedp.ByID),
		chromedp.SetValue(`//input[@name="departureDate"]`, "04-30-2018"),
		chromedp.SetValue(`//input[@name="returnDate"]`, "05-10-2018"),
		chromedp.WaitEnabled(`#cityOrigin`, chromedp.ByID),
		chromedp.WaitEnabled(`#cityDestination`, chromedp.ByID),
		chromedp.Sleep(2 * time.Second),
		chromedp.SendKeys(`#cityOrigin`, kb.ArrowDown+kb.ArrowDown+kb.ArrowDown+kb.ArrowDown, chromedp.ByID),
		chromedp.Sleep(4 * time.Second),
		chromedp.SendKeys(`#cityDestination`, kb.ArrowDown+kb.ArrowDown+kb.ArrowDown+kb.ArrowDown+kb.ArrowDown+kb.ArrowDown+kb.ArrowDown, chromedp.ByID),
		chromedp.Click(`//div[@class="form-group col-sm-3 text-right"]`),
		chromedp.WaitVisible(`//div[@class="dxp-flight-brand-price"]`),
		chromedp.Sleep(2 * time.Second),
		chromedp.Click(`//div[@class="dxp-flight-brand-price"]`),
		chromedp.Reload(),
		chromedp.WaitVisible(`//div[@class="dxp-flight-brand-price"]`),
		chromedp.Sleep(2 * time.Second),
		chromedp.Click(`//div[@class="dxp-flight-brand-price"]`),
		chromedp.Sleep(100 * time.Second),

		// chromedp.Text(`//*[@id="js-pjax-container"]/div[2]/div/div[2]/ul/li/p`, res),
	}
}

func submit(urlstr, sel, q string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel),
		chromedp.SendKeys(sel, q),
		chromedp.Submit(sel),
		chromedp.WaitNotPresent(`//*[@id="code_search"]/h2/svg`),
		chromedp.Text(`//*[@id="js-pjax-container"]/div[2]/div/div[2]/ul/li/p`, res),
	}
}
