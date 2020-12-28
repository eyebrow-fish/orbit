package store

import (
	"context"
	"database/sql"
	"fmt"
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

func ExecUnique(db *sql.DB, sql string, args ...interface{}) error {
	res, err := db.Exec(sql, args...)
	if err != nil {
		return err
	} else if rows, err := res.RowsAffected(); err != nil {
		return err
	} else if rows < 1 {
		return fmt.Errorf("could not create new row")
	}
	return nil
}

func Cleanup(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		panic(err)
	}
}
