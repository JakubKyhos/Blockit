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

	fmt.Println("-------------------------")
	fmt.Println("Whitelist reset was successfull!")
	fmt.Println("-------------------------")

	return nil
}

func DeleteWhitelistTempDom(db *database.Queries) error {
	err := db.DeleteWhitelistTempDoms(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete temp whitelisted domains: %w", err)
	}

	fmt.Println("-------------------------")
	fmt.Println("Deletion of temp domains was successfull!")
	fmt.Println("-------------------------")

	return nil
}

func DeleteWhitelistDom(db *database.Queries, name string) error {
	err := db.DeleteWhitelistDom(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't delete whitelisted domain: %w", err)
	}

	fmt.Println("-------------------------")
	fmt.Println("Deletion of whitelisted domain was successfull!")
	fmt.Println("-------------------------")

	return nil
}
