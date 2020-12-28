package main

import (
	"context"
	"database/sql"
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
		insert into Message(ChatId, Body, Timestamp) 
		select $1, $2, $3
		where not exists (
			select 1 from message where ChatId = $1 and Timestamp = $3
		)
		and exists (
			select 1 from Chat where Id = $1
		)
		`,
		req.ChatId,
		req.Body,
		postTime,
	)
	if err != nil {
		return nil, err
	}
	msg, err := store.QueryUniqueAndMap(
		db,
		"select * from Message where ChatId = $1 and Timestamp = $2",
		chat.Message{},
		req.ChatId,
		postTime,
	)
	if err != nil {
		return nil, err
	}
	return &post.ChatResp{Message: msg.(chat.Message)}, nil
}

func main() {
	ctx, err := store.DbCtx()
	if err != nil {
		panic(err)
	}
	lambda.StartWithContext(ctx, handle)
}
