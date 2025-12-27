package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/JakubKyhos/Blockit.git/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func CreateDomain(db *database.Queries, name string) error {
	var new_domain = database.CreateDomainParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	user, err := db.CreateDomain(context.Background(), new_domain)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				fmt.Printf("domain %s already exists\n", name)
				os.Exit(1)
			}
		}
		return err
	}

	fmt.Printf("user: %s was created\n", name)
	printDomain(user)
	return nil
}

func printDomain(domain database.Domain) {
	fmt.Printf(" * ID:      %v\n", domain.ID)
	fmt.Printf(" * Name:    %v\n", domain.Name)
}
