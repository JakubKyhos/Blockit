package whitelist

import (
	"context"
	"fmt"

	"github.com/JakubKyhos/Blockit.git/internal/database"
)

func ResetWhitelist(db *database.Queries) error {
	err := db.DeleteWhitelistDoms(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset whitelist: %w", err)
	}

	fmt.Println("Whitelist reset was successfull!")
	return nil
}

func DeleteWhitelistTempDom(db *database.Queries) error {
	err := db.DeleteWhitelistTempDoms(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete temp whitelisted domains: %w", err)
	}

	fmt.Println("Deletion of temp domains was successfull!")
	return nil
}

func DeleteWhitelistDom(db *database.Queries, name string) error {
	err := db.DeleteWhitelistDom(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't delete whitelisted domain: %w", err)
	}

	fmt.Println("Deletion of whitelisted domain was successfull!")
	return nil
}
