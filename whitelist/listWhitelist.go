package whitelist

import (
	"context"
	"fmt"

	"github.com/JakubKyhos/Blockit.git/internal/database"
)

func ListWhitelist(db *database.Queries) error {
	whitelist, err := db.GetWhitelistDoms(context.Background())
	if err != nil {
		return fmt.Errorf("failed to list domains: %v", err)
	}

	if len(whitelist) == 0 {
		return fmt.Errorf("whitelist is empty use 'add whitelist webpage' to populate whitelist")
	}

	fmt.Println("-------------------------")
	for _, okweb := range whitelist {
		PrintWhitelist(okweb)
	}

	return nil
}
