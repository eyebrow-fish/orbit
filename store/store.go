package store

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

func DbCtx() (context.Context, error) {
	root := context.Background()
	db, err := sql.Open("postgres", "user=user dbname=orbit sslmode=verify-full")
	if err != nil {
		return nil, err
	}
	return context.WithValue(root, "db", db), nil
}
