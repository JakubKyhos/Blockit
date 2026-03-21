package blacklist

import (
	"context"
	"fmt"

	"github.com/JakubKyhos/Blockit.git/internal/database"
)

func ResetBlacklist(db *database.Queries) error {
	err := db.DeleteBlacklistDoms(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset blacklist: %w", err)
	}

	fmt.Println("-------------------------")
	fmt.Println("Blacklist reset was successfull!")
	fmt.Println("-------------------------")

	return nil
}

func DeleteBlacklistDom(db *database.Queries, name string) error {
	err := db.DeleteBlacklistDom(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't delete blacklisted domain: %w", err)
	}

	fmt.Println("-------------------------")
	fmt.Println("Deletion of blacklisted domain was successfull!")
	fmt.Println("-------------------------")

	return nil
}
