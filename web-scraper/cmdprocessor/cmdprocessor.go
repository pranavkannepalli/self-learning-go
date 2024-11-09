package cmdprocessor

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pranavkannepalli/self-learning-go/web-scraper/scraper"
)

func ProcessInput() *scraper.ScraperArgs {
	var input string
	fmt.Print("Pick a website to scrape or type X to escape: ")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		return ProcessInput()
	}
	if input == "X" {
		return nil
	}
	split := strings.Fields(input)
	args := scraper.ScraperArgs{Website: split[0], Verbose: true}

	for idx, val := range split {
		if val == "--d" {
			if len(split) < idx+1 {
				return ProcessInput()
			}
			depth, _ := strconv.Atoi(split[idx+1])
			args.Depth = depth
		}
		if val == "--s" && args.Verbose {
			args.Verbose = false
		}
	}

	return &args
}
