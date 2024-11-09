package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type ScraperArgs struct {
	Website string
	Depth   int
	Verbose bool
}

func RecurseScrape(args *ScraperArgs) int {
	fmt.Println(args.Depth)
	depth := args.Depth
	if depth == 0 {
		depth = 2
	}
	website := args.Website
	if !strings.HasPrefix(website, "https://") {
		website = "https://" + website
	}

	c := colly.NewCollector()
	c.MaxDepth = depth

	numVisited := 0
	numFound := 0

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
		numFound++
	})

	c.OnRequest(func(r *colly.Request) {
		if args.Verbose {
			fmt.Println("Visiting", r.URL)
		}
		numVisited++
	})

	c.Visit(website)

	return numVisited
}
