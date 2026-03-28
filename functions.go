package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() []string {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		return nil
	}
	line := scanner.Text()
	line = strings.TrimSpace(line)
	return strings.Fields(line)
}

func help() {
	fmt.Println(`-------------------------
- help: shows available commands

- setup: loads known TLD's into DB (all default to false, meaning unblocked)

- add {whiteliset/whitelisttemp/blacklist} {domain}: adds domain to corresponding DB, i.e. "add blacklist example.com" (whitelisttemp lasts one hour)

- delete {whitelist/blacklist} {domain}: deletes domain from corresponding DB (you can do "delete whitelist temp" to delete expired whitelisted domains)

- list {whitelist/blacklist/tld}: show all domains/TLD's in corresponding DB

- reset {whitelist/blacklist/tld}: delete all from corresponding DB

- blockstate {true/false} {tld/global}: change whether TLD is blocked or not - example "blockstate true net" ("global" changes blockstate for all TLD's)

	WARNING! CHANGING BLOCKSTATE TO TRUE FOR TLD's LIKE com, org, net, io, etc. CAN CAUSE SOME WEBSITES TO MULFUNCTION IF THEY USE MULTIPLE DOMAINS (LIKE PAYMENT GATES).

	EXAMPLE: you block "com" and whitelist "reddit.com", since Reddit uses other domains such as "reddit-image.s3.amazonaws.com", it won't work as you might expect.

- quit: close Blockit CLI
-------------------------`)
}
