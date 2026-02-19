package TLDs

import (
	"fmt"
	"strings"

	"github.com/JakubKyhos/Blockit.git/internal/database"
	"github.com/JakubKyhos/Blockit.git/scraper"
)

func Setup(dbQueries *database.Queries, webPage string) error {
	TopLevelDomains, err := scraper.CrawlPage(webPage)
	if err != nil {
		return fmt.Errorf("failed to fetch domains: %s\n", err)
	}

	parsedTLD := ParseTLDList(TopLevelDomains)
	for i := 0; i < len(parsedTLD); i++ {
		err = SetupDomain(dbQueries, parsedTLD[i])
		if err != nil {
			return fmt.Errorf("failed to create domain for DB: %v", err)
		}
	}
	return nil
}

func ParseTLDList(text string) []string {
	// Split into lines
	raw := strings.Split(text, "\n")

	var tlds []string
	for _, line := range raw {
		line = strings.TrimSpace(line)
		if line == "" {
			continue // skip empty lines
		}
		if strings.HasPrefix(line, "#") {
			continue // skip comment header lines from IANA
		}
		tlds = append(tlds, line)
	}

	return tlds
}
