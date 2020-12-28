package testutil

import (
	"context"
	"github.com/eyebrow-fish/orbit/store"
	"testing"
	"zombiezen.com/go/postgrestest"
)

func DbCtx(t *testing.T) (context.Context, error) {
	srvCtx := context.Background()
	srv, err := postgrestest.Start(srvCtx)
	if err != nil {
		return nil, err
	}
	t.Cleanup(srv.Cleanup)
	db, err := srv.NewDatabase(srvCtx)
	if err != nil {
		return nil, err
	}
	ctx := context.WithValue(context.Background(), "db", &store.Db{DB: db})
	return ctx, nil
}
