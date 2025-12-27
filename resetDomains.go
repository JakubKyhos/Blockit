package main

import (
	"context"
	"fmt"

	"github.com/JakubKyhos/Blockit.git/internal/database"
)

func resetDomains(db *database.Queries) error {
	err := db.DeleteDomains(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}

	fmt.Println("Database reset successfully!")
	return nil
}
