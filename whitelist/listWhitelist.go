package whitelist

import (
	"context"
	"fmt"

	"github.com/JakubKyhos/Blockit.git/internal/database"
)

func ListWhitelist(db *database.Queries) ([]database.Whitelist, error) {
	whitelist, err := db.GetWhitelistDoms(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to list domains: %v", err)
	}

	var webPageList []database.Whitelist
	for i := 0; i < len(whitelist); i++ {
		page := database.Whitelist{
			ID:        whitelist[i].ID,
			CreatedAt: whitelist[i].CreatedAt,
			ExpiresAt: whitelist[i].ExpiresAt,
			Name:      whitelist[i].Name,
		}
		webPageList = append(webPageList, page)
	}
	return webPageList, nil
}
