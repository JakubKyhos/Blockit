package whitelist

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/JakubKyhos/Blockit.git/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func CreateWhitelist(db *database.Queries, name string) error {
	var new_whitelist = database.CreateWhitelistDomParams{
		ID:   uuid.New(),
		Name: strings.ToLower(strings.TrimSpace(name)),
	}

	whitelist, err := db.CreateWhitelistDom(context.Background(), new_whitelist)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				fmt.Printf("domain %s already exists in database\n", name)
				os.Exit(1)
			}
		}
		return err
	}

	PrintWhitelist(whitelist)

	return nil
}

func CreateWhitelistTemp(db *database.Queries, name string) error {
	currentTime := time.Now().Local()
	duration := currentTime.Add(time.Hour * 1)
	nullTime := sql.NullTime{
		Time:  duration,
		Valid: true,
	}

	var new_whitelistTemp = database.CreateWhitelistDomTempParams{
		ID:        uuid.New(),
		ExpiresAt: nullTime,
		Name:      strings.ToLower(name),
	}

	whitelist, err := db.CreateWhitelistDomTemp(context.Background(), new_whitelistTemp)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				fmt.Printf("tempDomain %s already exists in database\n", name)
				os.Exit(1)
			}
		}
		return err
	}

	fmt.Println("-------------------------")
	PrintWhitelist(whitelist)

	return nil
}

func PrintWhitelist(whitelist database.Whitelist) {
	fmt.Printf(" * %v has been added to whitelist successfully\n", whitelist.Name)
	fmt.Printf(" * ID:      	 %v\n", whitelist.ID)
	fmt.Printf(" * CreatedAt 	 %s\n", whitelist.CreatedAt)
	fmt.Printf(" * ExpiresAt 	 %s\n", whitelist.ExpiresAt.Time)
	fmt.Printf(" * Name:    	 %s\n", whitelist.Name)
	fmt.Println("-------------------------")
}
