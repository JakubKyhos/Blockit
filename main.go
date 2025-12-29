package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/JakubKyhos/Blockit.git/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	const webPage = "https://data.iana.org/TLD/tlds-alpha-by-domain.txt"

	for {
		words := getInput()
		if len(words) == 0 {
			continue
		}
		switch words[0] {
		case "setup":
			err = setup(dbQueries, webPage)
			if err != nil {
				fmt.Printf("failed to set up DB: %v", err)
				continue
			}
			fmt.Println("-------------------------")
			fmt.Println("DB has been set up successfully")
			fmt.Println("-------------------------")
		case "list":
			dblist, err := listDomains(dbQueries)
			if err != nil {
				fmt.Printf("failed to list DB: %v\n", err)
				continue
			}
			if len(dblist) == 0 {
				fmt.Println("DB is empty use 'setup' to populate DB")
				continue
			}
			fmt.Printf("%v\n", dblist)
			for _, domain := range dblist {
				fmt.Println("-------------------------")
				printDomain(domain)
				fmt.Println("-------------------------")
			}
		case "reset":
			err = resetDomains(dbQueries)
			if err != nil {
				fmt.Printf("failed to reset DB: %v", err)
				continue
			}
		case "blockstate":
			if len(words) != 3 {
				fmt.Println("blockstate needs 3 arguments: blockstate true/false domainName/global")
				continue
			}
			err := blockStatewrapper(dbQueries, words[1], words[2])
			if err != nil {
				fmt.Printf("there was an issue with block state cmd: %v\n", err)
				continue
			}

		case "quit":
			fmt.Println("quiting BlockIt settings")
			return
		default:
			fmt.Println("unknown command")
		}
	}
}
