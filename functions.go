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

func parseTLDList(text string) []string {
	// Split into lines
	raw := strings.Split(text, "\n")

	var tlds []string
	for _, line := range raw {
		line = strings.TrimSpace(line)
		if line == "" {
			continue // skip empty lines
		}
		if strings.HasPrefix(line, "#") {
			continue // skip comment header lines from IANA
		}
		tlds = append(tlds, line)
	}

	return tlds
}
