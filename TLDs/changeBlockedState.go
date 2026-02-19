package TLDs

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/JakubKyhos/Blockit.git/internal/database"
)

func ChangeBlockedState(db *database.Queries, args database.DomainBlockStateParams) (database.Domain, error) {
	domain, err := db.DomainBlockState(context.Background(), args)
	if err != nil {
		return database.Domain{}, fmt.Errorf("failed to update %s's blocked state: %v", args.Name, err)
	}

	fmt.Println("blocked state was changed sucessfully")
	return domain, nil
}

func ChangeBlockedStateGlobal(db *database.Queries, isBlocked bool) ([]database.Domain, error) {
	domains, err := db.DomainsBlockedStateGlobal(context.Background(), isBlocked)
	if err != nil {
		return nil, fmt.Errorf("failed to update blocked state globaly: %v", err)
	}

	fmt.Println("global blocked state was changed successfully")
	return domains, nil
}

func BlockStatewrapper(db *database.Queries, boolean string, subject string) error {
	state, err := strconv.ParseBool(strings.TrimSpace(strings.ToLower(boolean)))
	if err != nil {
		return fmt.Errorf("failed toconvert bool value: %v", err)
	}
	subject = strings.TrimSpace(strings.ToLower(subject))
	var args = database.DomainBlockStateParams{
		IsBlocked: state,
		Name:      subject,
	}
	if subject == "global" {
		datBase, err := ChangeBlockedStateGlobal(db, state)
		if err != nil {
			return fmt.Errorf("failed to change blockstate globaly: %v", err)
		}
		for _, domain := range datBase {
			fmt.Println("-------------------------")
			PrintDomain(domain)
			fmt.Println("-------------------------")
		}
		return nil
	} else {
		domain, err := ChangeBlockedState(db, args)
		if err != nil {
			return fmt.Errorf("failed to change block state for domain %s: %v", domain.Name, err)
		}
		fmt.Println("-------------------------")
		PrintDomain(domain)
		fmt.Println("-------------------------")
		return nil
	}
}
