package services

// import (
// 	"bytes"
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"path/filepath"

// 	"github.com/chromedp/chromedp"
// )

import (
	"context"
	"fmt"
	"os"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

const (
	websiteUrl = "https://www.detik.com/"
	dockerUrl  = "wss://localhost:9222"
)

func CrawlDetik() {
	// create allocator context for use with creating a browser context later
	allocatorContext, cancel := chromedp.NewRemoteAllocator(context.Background(), dockerUrl)
	defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(websiteUrl),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			res, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),
	)

	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("output.html", []byte(res), 0644)

	if err != nil {
		fmt.Println(err)
	}

	// resp, err := chromedp.RunResponse(ctx, chromedp.Navigate(websiteUrl))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("status code:", resp.Status)
}
