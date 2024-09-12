package dbexample

import (
	"context"
	"coree/ent"
	"coree/ent/user"
	"fmt"
	"log"
)

var dbClient *ent.Client

func queryUser(ctx context.Context, client *ent.Client, name string) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name(name)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func createUser(ctx context.Context, client *ent.Client, name string, age int) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}
