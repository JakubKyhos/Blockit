package whitelist

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/JakubKyhos/Blockit.git/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func CreateWhitelist(db *database.Queries, name string) error {
	var new_whitelist = database.CreateWhitelistDomParams{
		ID:   uuid.New(),
		Name: name,
	}

	whitelist, err := db.CreateWhitelistDom(context.Background(), new_whitelist)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				fmt.Printf("domain %s is already whitelisted\n", name)
				os.Exit(1)
			}
		}
		return err
	}

	fmt.Printf("domain: %s was whitelisted\n", name)
	PrintWhitelist(whitelist)
	return nil
}

func CreateWhitelistTemp(db *database.Queries, name string) error {
	currentTime := time.Now().Local()
	duration := currentTime.Add(time.Minute * 1)
	nullTime := sql.NullTime{
		Time:  duration,
		Valid: true,
	}

	var new_whitelistTemp = database.CreateWhitelistDomTempParams{
		ID:        uuid.New(),
		ExpiresAt: nullTime,
		Name:      name,
	}

	whitelist, err := db.CreateWhitelistDomTemp(context.Background(), new_whitelistTemp)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				fmt.Printf("tempDomain %s already exists\n", name)
				os.Exit(1)
			}
		}
		return err
	}

	fmt.Printf("tempDom: %s was created\n", name)
	PrintWhitelist(whitelist)
	return nil
}

func PrintWhitelist(whitelist database.Whitelist) {
	fmt.Printf(" * ID:      	 %v\n", whitelist.ID)
	fmt.Printf(" * CreatedAt 	 %s\n", whitelist.CreatedAt)
	fmt.Printf(" * ExpiresAt 	 %s\n", whitelist.ExpiresAt.Time)
	fmt.Printf(" * Name:    	 %s\n", whitelist.Name)
}
