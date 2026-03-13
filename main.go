package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/JakubKyhos/Blockit.git/TLDs"
	"github.com/JakubKyhos/Blockit.git/internal/database"
	"github.com/JakubKyhos/Blockit.git/whitelist"
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
			err = TLDs.Setup(dbQueries, webPage)
			if err != nil {
				fmt.Printf("failed to set up DB: %v", err)
				continue
			}
			fmt.Println("-------------------------")
			fmt.Println("DB has been set up successfully")
			fmt.Println("-------------------------")
		case "add":
			if len(words) != 3 {
				fmt.Println("add needs three arguments: add whitelist/whitelisttemp domainName")
				continue
			}
			switch words[1] {
			case "whitelist":
				err = whitelist.CreateWhitelist(dbQueries, words[2])
				if err != nil {
					fmt.Printf("failed to add %v to DB: %v", words[2], err)
					continue
				}
				fmt.Println("-------------------------")
				fmt.Printf("%v has been added to whitelist successfully\n", words[2])
				fmt.Println("-------------------------")
			case "whitelisttemp":
				err = whitelist.CreateWhitelistTemp(dbQueries, words[2])
				if err != nil {
					fmt.Printf("failed to add %v to DB: %v", words[2], err)
					continue
				}
				fmt.Println("-------------------------")
				fmt.Printf("%v has been added temporarily to whitelist\n", words[2])
				fmt.Println("-------------------------")
			default:
				fmt.Println("Unknown argument. Use whitelist or whitelisttemp.")
			}
		case "delete":
			if len(words) != 3 {
				fmt.Println("delete needs three arguments: delete whitelist domainName/temp")
				continue
			}
			switch words[1] {
			case "whitelist":
				if words[2] == "temp" {
					err = whitelist.DeleteWhitelistTempDom(dbQueries)
					if err != nil {
						fmt.Printf("failed to delete old temporary domains from whitelist: %v", err)
						continue
					}
				} else {
					err = whitelist.DeleteWhitelistDom(dbQueries, words[2])
					if err != nil {
						fmt.Printf("failed to delete domain %v from whitelist: %v", words[2], err)
						continue
					}
				}
			default:
				fmt.Println("Unknown argument. Use whitelist.")
			}
		case "list":
			if len(words) != 2 {
				fmt.Println("list needs two arguments: list tld/whitelist")
				continue
			}
			switch words[1] {
			case "whitelist":
				dblist, err := whitelist.ListWhitelist(dbQueries)
				if err != nil {
					fmt.Printf("failed to list whitelist from DB: %v\n", err)
					continue
				}
				if len(dblist) == 0 {
					fmt.Println("whitelist is empty use 'add whitelist webpage' to populate whitelist")
					continue
				}
				for _, okweb := range dblist {
					fmt.Println("-------------------------")
					whitelist.PrintWhitelist(okweb)
					fmt.Println("-------------------------")
				}
			case "tld":
				dblist, err := TLDs.ListDomains(dbQueries)
				if err != nil {
					fmt.Printf("failed to list tld's from DB: %v\n", err)
					continue
				}
				if len(dblist) == 0 {
					fmt.Println("tld's DB is empty, use 'setup' to populate it")
					continue
				}
				fmt.Printf("%v\n", dblist)
				for _, domain := range dblist {
					fmt.Println("-------------------------")
					TLDs.PrintDomain(domain)
					fmt.Println("-------------------------")
				}
			default:
				fmt.Println("Unknown argument. Use tld or whitelist.")
			}

		case "reset":
			if len(words) != 2 {
				fmt.Println("reset requires two arguments: reset tld/whitelist")
			}
			switch words[1] {
			case "tld":
				err = TLDs.ResetDomains(dbQueries)
				if err != nil {
					fmt.Printf("failed to reset DB: %v", err)
					continue
				}
			case "whitelist":
				err = whitelist.ResetWhitelist(dbQueries)
				if err != nil {
					fmt.Printf("failed to reset DB: %v", err)
					continue
				}
			default:
				fmt.Println("Unknown argument. Use tld or whitelist.")
			}

		case "blockstate":
			if len(words) != 3 {
				fmt.Println("blockstate needs 3 arguments: blockstate true/false domainName/global")
				continue
			}
			err := TLDs.BlockStatewrapper(dbQueries, words[1], words[2])
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
