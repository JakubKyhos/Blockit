package main

import (
	"fmt"

	"github.com/JakubKyhos/Blockit.git/scraper"
)

func main() {
	const webPage = "https://data.iana.org/TLD/tlds-alpha-by-domain.txt"

	for {
		words := getInput()
		if len(words) == 0 {
			continue
		}
		switch words[0] {
		case "crawl":
			TopLevelDomains, err := scraper.CrawlPage(webPage)
			if err != nil {
				fmt.Printf("failed to fetch domains: %s\n", err)
				continue
			}
			parsedTLD := parseTLDList(TopLevelDomains)
			fmt.Printf("%s\n", parsedTLD)
		case "quit":
			fmt.Println("quiting BlockIt settings")
			return
		default:
			fmt.Println("unknown command")
		}
	}
}
