package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/eyebrow-fish/orbit/chat/post"
	"github.com/eyebrow-fish/orbit/store"
)

func main() {
	ctx, err := store.DbCtx()
	if err != nil {
		panic(err)
	}
	lambda.StartWithContext(ctx, post.Handle)
}
