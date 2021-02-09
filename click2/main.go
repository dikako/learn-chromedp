package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	//create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://pkg.go.dev/time`),

		// wait for footer element is visible (ie. page is loaded)
		chromedp.WaitVisible(`body > footer`),

		// find and click "Example" link
		chromedp.Click(`#example-After`, chromedp.NodeVisible),

		// retrive the text pf the <pre>
		chromedp.Text(`#example-After pre.Documentation-exampleCode`, &example),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Go's time.After example:\n%s", example)
}
