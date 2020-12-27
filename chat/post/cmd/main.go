package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/eyebrow-fish/orbit/chat/post"
	"github.com/eyebrow-fish/orbit/store"
)

func handle(ctx context.Context, req post.ChatReq) (*post.ChatResp, error) {
	return nil, nil
}

func main() {
	ctx, err := store.DbCtx()
	if err != nil {
		panic(err)
	}
	lambda.StartWithContext(ctx, handle)
}
