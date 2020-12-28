package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/eyebrow-fish/orbit/chat"
	"github.com/eyebrow-fish/orbit/chat/create"
	"github.com/eyebrow-fish/orbit/store"
)

func handle(ctx context.Context, req create.ChatReq) (*create.ChatResp, error) {
	db := ctx.Value("db").(*store.Db)
	err := db.ExecUnique(
		`
		insert into Chat(Name)
		select $1
		where not exists (
			select 1 from Chat where Name = $1
		)
		`,
		req.Name,
	)
	if err != nil {
		return nil, err
	}
	resp, err := db.QueryUnique(chat.Chat{}, "select * from Chat where Name = $1", req.Name)
	return &create.ChatResp{Chat: resp.(chat.Chat)}, nil
}

func main() {
	ctx, err := store.DbCtx()
	if err != nil {
		panic(err)
	}
	lambda.StartWithContext(ctx, handle)
}
