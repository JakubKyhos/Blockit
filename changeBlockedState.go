package main

import (
	"context"
	"fmt"

	"github.com/JakubKyhos/Blockit.git/internal/database"
)

func changeBlockedState(db *database.Queries, args database.DomainBlockStateParams) (database.Domain, error) {
	domain, err := db.DomainBlockState(context.Background(), args)
	if err != nil {
		return database.Domain{}, fmt.Errorf("failed to update %s's blocked state: %v", args.Name, err)
	}
	return domain, nil
}

func changeBlockedStateGlobal(db *database.Queries, isBlocked bool) ([]database.Domain, error) {
	domains, err := db.DomainsBlockedStateGlobal(context.Background(), isBlocked)
	if err != nil {
		return nil, fmt.Errorf("failed to update blocked state globaly: %v", err)
	}

	return domains, nil
}
