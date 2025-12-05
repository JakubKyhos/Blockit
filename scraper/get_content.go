package scraper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	// The file contents are inside <pre>
	pre := doc.Find("pre").First().Text()
	if pre != "" {
		return strings.TrimSpace(pre)
	}

	// fallback: whole body text
	body := doc.Find("body").First().Text()
	return strings.TrimSpace(body)
}
