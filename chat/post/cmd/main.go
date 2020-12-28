package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/eyebrow-fish/orbit/chat"
	"github.com/eyebrow-fish/orbit/chat/post"
	"github.com/eyebrow-fish/orbit/store"
	"time"
)

func handle(ctx context.Context, req post.ChatReq) (*post.ChatResp, error) {
	postTime := time.Now().UnixNano()
	db := ctx.Value("db").(*sql.DB)
	err := store.ExecUnique(
		db,
		`
		insert into message(chatId, body, timestamp) 
		select $1, $2, $3
		where not exists (
			select 1 from message where chatId = $1 and timestamp = $3
		)
		`,
		req.ChatId,
		req.Body,
		postTime,
	)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from message where chatId = $1 and timestamp = $2", req.ChatId, postTime)
	if err != nil {
		return nil, err
	}
	defer store.Cleanup(rows)
	if !rows.Next() {
		return nil, fmt.Errorf("could not insert data")
	}
	var (
		id        int
		chatId    int
		body      string
		timestamp int64
	)
	err = rows.Scan(&id, &chatId, &body, &timestamp)
	if err != nil {
		return nil, err
	}
	return &post.ChatResp{Message: chat.Message{ChatId: chatId, Timestamp: timestamp, Body: body}}, nil
}

func main() {
	ctx, err := store.DbCtx()
	if err != nil {
		panic(err)
	}
	lambda.StartWithContext(ctx, handle)
}
