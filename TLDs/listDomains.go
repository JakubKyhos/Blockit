package TLDs

import (
	"context"
	"fmt"

	"github.com/JakubKyhos/Blockit.git/internal/database"
)

func ListDomains(db *database.Queries) ([]database.Domain, error) {
	domains, err := db.GetDomains(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to list domains: %v", err)
	}

	var domainList []database.Domain
	for i := 0; i < len(domains); i++ {
		domain := database.Domain{
			ID:        domains[i].ID,
			CreatedAt: domains[i].CreatedAt,
			UpdatedAt: domains[i].UpdatedAt,
			Name:      domains[i].Name,
			IsBlocked: domains[i].IsBlocked,
		}
		domainList = append(domainList, domain)
	}
	return domainList, nil
}
