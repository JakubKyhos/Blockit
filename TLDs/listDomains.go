package TLDs

import (
	"context"
	"fmt"

	"github.com/JakubKyhos/Blockit.git/internal/database"
)

func ListDomains(db *database.Queries) error {
	domains, err := db.GetDomains(context.Background())
	if err != nil {
		return fmt.Errorf("failed to list tld's from DB: %v", err)
	}

	if len(domains) == 0 {
		return fmt.Errorf("tld's DB is empty, use 'setup' to populate it")
	}

	fmt.Println("-------------------------")
	for _, domain := range domains {
		PrintDomain(domain)
	}

	return nil
}
