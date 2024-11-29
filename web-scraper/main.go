package main

import (
	"fmt"

	"github.com/pranavkannepalli/self-learning-go/web-scraper/cmdprocessor"
	"github.com/pranavkannepalli/self-learning-go/web-scraper/scraper"
)

// TODO: add logger system

func main() {
	fmt.Println("Starting webscraper")

	for {
		args := cmdprocessor.ProcessInput()
		if args == nil {
			break
		}
		visited := scraper.RecurseScrape(args)
		fmt.Printf("Visited %v links.\n", visited)
	}
}
