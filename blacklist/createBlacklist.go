package blacklist

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/JakubKyhos/Blockit.git/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func CreateBlacklist(db *database.Queries, name string) error {
	var new_blacklist = database.CreateBlacklistDomParams{
		ID:   uuid.New(),
		Name: strings.ToLower(name),
	}

	blacklist, err := db.CreateBlacklistDom(context.Background(), new_blacklist)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				fmt.Printf("domain %s is already in database\n", name)
				os.Exit(1)
			}
		}
		return err
	}

	fmt.Printf("domain: %s was blacklisted\n", name)
	PrintBlacklist(blacklist)
	return nil
}

func PrintBlacklist(blacklist database.Blacklist) {
	fmt.Printf(" * ID:      	 %v\n", blacklist.ID)
	fmt.Printf(" * CreatedAt 	 %s\n", blacklist.CreatedAt)
	fmt.Printf(" * Name:    	 %s\n", blacklist.Name)
}
