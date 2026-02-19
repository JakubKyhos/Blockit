package TLDs

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/JakubKyhos/Blockit.git/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func SetupDomain(db *database.Queries, name string) error {
	var new_domain = database.CreateDomainParams{
		ID:   uuid.New(),
		Name: strings.ToLower(name),
	}

	domain, err := db.CreateDomain(context.Background(), new_domain)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				fmt.Printf("domain %s already exists\n", name)
				os.Exit(1)
			}
		}
		return err
	}

	fmt.Printf("domain: %s was created\n", name)
	PrintDomain(domain)
	return nil
}

func PrintDomain(domain database.Domain) {
	fmt.Printf(" * ID:      	 %v\n", domain.ID)
	fmt.Printf(" * CreatedAt 	 %s\n", domain.CreatedAt)
	fmt.Printf(" * UpdatedAt 	 %s\n", domain.UpdatedAt)
	fmt.Printf(" * Name:    	 %s\n", domain.Name)
	fmt.Printf(" * IsBlocked:    %t\n", domain.IsBlocked)
}
