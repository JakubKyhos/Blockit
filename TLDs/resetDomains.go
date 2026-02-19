package TLDs

import (
	"context"
	"fmt"

	"github.com/JakubKyhos/Blockit.git/internal/database"
)

func ResetDomains(db *database.Queries) error {
	err := db.DeleteDomains(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete TLDs: %w", err)
	}

	fmt.Println("TLD database reset was successfull!")
	return nil
}
