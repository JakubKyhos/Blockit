package scraper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	h1 := doc.Find("h1").First().Text()
	return strings.TrimSpace(h1)
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	HTMLbody := doc.Find("body")
	var pre string
	if HTMLbody.Length() > 0 {
		pre = HTMLbody.Find("pre").First().Text()
	} else {
		pre = doc.Find("pre").First().Text()
	}

	return strings.TrimSpace(pre)
}
