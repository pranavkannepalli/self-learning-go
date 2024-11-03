package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("Starting webscraper")

	c := colly.NewCollector()
	c.MaxDepth = 2

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	var website string
	fmt.Print("Pick a website to scrape: ")
	fmt.Scanln(&website)
	c.Visit(website)
}
