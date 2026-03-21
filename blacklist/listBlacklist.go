package blacklist

import (
	"context"
	"fmt"

	"github.com/JakubKyhos/Blockit.git/internal/database"
)

func ListBlacklist(db *database.Queries) error {
	blacklist, err := db.GetBlacklistDoms(context.Background())
	if err != nil {
		return fmt.Errorf("failed to list domains: %v", err)
	}

	if len(blacklist) == 0 {
		return fmt.Errorf("blacklist is empty use 'add blacklist webpage' to populate blacklist")
	}

	fmt.Println("-------------------------")
	for _, blocked := range blacklist {
		PrintBlacklist(blocked)
	}

	return nil
}
