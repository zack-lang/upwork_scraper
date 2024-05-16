package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	url := "https://www.upwork.com/freelance-jobs/golang/"
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"),
	)

	// Find and print all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Printf("Element: %s,\n Link: %s\n", e.Text, e.Attr("href"))
		fmt.Println("---------------------------------------------------------------")
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "*/**")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.5")
		r.Headers.Set("DNT", "1")
		r.Headers.Set("Alt-Used", "www.upwork.com")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
		r.Headers.Set("Sec-Fetch-Dest", "document")
		r.Headers.Set("Sec-Fetch-Mode", "navigate")
		r.Headers.Set("Sec-Fetch-Site", "none")
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		if r.StatusCode == http.StatusForbidden {
			fmt.Println("Request URL:", r.Request.URL, "failed with error:", err)
			// Wait
			time.Sleep(15 * time.Second)
			// Retry with a different user agent
			r.Request.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
			fmt.Println("---------------------------------------------------------------")
			c.Visit(url)
			fmt.Println("****************************************************************")
		}
	})

	c.Visit(url)
}
