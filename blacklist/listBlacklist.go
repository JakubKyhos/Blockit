package blacklist

import (
	"context"
	"fmt"

	"github.com/JakubKyhos/Blockit.git/internal/database"
)

func ListBlacklist(db *database.Queries) ([]database.Blacklist, error) {
	blacklist, err := db.GetBlacklistDoms(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to list domains: %v", err)
	}

	var domainList []database.Blacklist
	for i := 0; i < len(blacklist); i++ {
		domain := database.Blacklist{
			ID:        blacklist[i].ID,
			CreatedAt: blacklist[i].CreatedAt,
			Name:      blacklist[i].Name,
		}
		domainList = append(domainList, domain)
	}
	return domainList, nil
}
