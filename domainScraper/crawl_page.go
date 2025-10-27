package scraper

import (
	"fmt"
)

func crawlPage(rawCurrentURL string) (PageData, error) {
	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return PageData{
			URL:            rawCurrentURL,
			H1:             "",
			FirstParagraph: "",
		}, err
	}

	// Extract all the data we care about and store it
	pageData := extractPageData(htmlBody, rawCurrentURL)

	return pageData, nil
}
