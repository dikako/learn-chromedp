package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

func main() {
	// Create Context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Run Test with emulate using device & viewport!

	var ss1, ss2, ss3 []byte
	if err := chromedp.Run(ctx,

		// Emulate Iphone 11 ProMax Potrait
		chromedp.Emulate(device.IPhone11ProMax),
		chromedp.Navigate(`https://dikakoko.medium.com/how-to-make-api-automation-testing-with-cucumber-ruby-and-httparty-461105963389`),
		chromedp.CaptureScreenshot(&ss1),

		// Reset emulate
		chromedp.Emulate(device.Reset),

		// Emulate Iphone 11 ProMax Landscape
		chromedp.Emulate(device.IPhone11ProMaxlandscape),
		chromedp.Navigate(`https://dikakoko.medium.com/how-to-make-api-automation-testing-with-cucumber-ruby-and-httparty-461105963389`),
		chromedp.CaptureScreenshot(&ss2), // for screenshoot result

		// Emulate with viewport
		chromedp.EmulateViewport(1000, 1000),
		chromedp.Navigate(`https://dikakoko.medium.com/how-to-make-api-automation-testing-with-cucumber-ruby-and-httparty-461105963389`),
		chromedp.CaptureScreenshot(&ss3),
	); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("screenshot/iphone_11_potrait.png", ss1, 0o644); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("screenshot/iphone_11_landscape.png", ss2, 0o644); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("screenshot/viewport_1000x1000.png", ss3, 0o644); err != nil {
		log.Fatal(err)
	}

	log.Printf("Success Test Emulate with Screenshoot result")
}

// Thankyou!!
