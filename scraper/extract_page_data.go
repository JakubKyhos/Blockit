package scraper

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
}

func extractPageData(html, pageURL string) PageData {
	h1 := getH1FromHTML(html)
	firstParagraph := getFirstParagraphFromHTML(html)

	return PageData{
		URL:            pageURL,
		H1:             h1,
		FirstParagraph: firstParagraph,
	}
}
