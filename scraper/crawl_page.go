package scraper

import (
	"fmt"
)

func CrawlPage(rawCurrentURL string) (string, error) {
	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return "", err
	}

	// Extract all the data we care about and store it
	topLevelDomains := getFirstParagraphFromHTML(htmlBody)

	return topLevelDomains, nil
}
