package main

import (
	"fmt"

	"github.com/JakubKyhos/Blockit.git/internal/database"
	"github.com/JakubKyhos/Blockit.git/scraper"
)

func setup(dbQueries *database.Queries, webPage string) error {
	TopLevelDomains, err := scraper.CrawlPage(webPage)
	if err != nil {
		return fmt.Errorf("failed to fetch domains: %s\n", err)
	}

	parsedTLD := parseTLDList(TopLevelDomains)
	for i := 0; i < len(parsedTLD); i++ {
		err = SetupDomain(dbQueries, parsedTLD[i])
		if err != nil {
			return fmt.Errorf("failed to create domain for DB: %v", err)
		}
	}
	return nil
}
