package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/eyebrow-fish/orbit/chat"
	"github.com/eyebrow-fish/orbit/chat/create"
)

func handle(ctx context.Context, req create.ChatReq) (*create.ChatResp, error) {
	db := ctx.Value("db").(*sql.DB)
	res, err := db.Exec(`
			insert into chat(name)
			select $1
			where not exists (
			    select 1 from chat where name = $1
			)
		`,
		req.Name,
	)
	if err != nil {
		return nil, err
	} else if rows, err := res.RowsAffected(); err != nil {
		return nil, err
	} else if rows < 1 {
		return nil, fmt.Errorf("could not create new chat")
	}
	rows, err := db.Query("select * from chat where name = $1", req.Name)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			panic(err)
		}
	}()
	if !rows.Next() {
		return nil, fmt.Errorf("could not insert data")
	}
	var (
		id   int
		name string
	)
	if err = rows.Scan(&id, &name); err != nil {
		return nil, err
	}
	return &create.ChatResp{Chat: chat.Chat{Id: id, Name: name}}, nil
}

func main() {
	lambda.Start(handle)
}
