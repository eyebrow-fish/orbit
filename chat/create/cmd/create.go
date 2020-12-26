package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/eyebrow-fish/orbit/chat/create"
)

func handle(ctx context.Context, req create.ChatReq) (*create.ChatResp, error) {
	db := ctx.Value("db").(*sql.DB)
	res, err := db.Exec(`insert into chat(id, name) values (0, ?)`, req.Name)
	if err != nil {
		return nil, err
	}
	if rows, err := res.RowsAffected(); err != nil {
		return nil, err
	} else if rows < 1 {
		return nil, fmt.Errorf("no rows inserted for %v", req)
	}
	return &create.ChatResp{Name: req.Name}, nil
}

func main() {
	lambda.Start(handle)
}
