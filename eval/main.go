package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	// Create Context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res []string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.google.com/`),
		chromedp.WaitVisible(`#main`, chromedp.ByID),
		chromedp.Evaluate(`Object.keys(window);`, &res),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("windows object keys: %v", res)
}
