package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/eyebrow-fish/orbit/chat/create"
)

func handle(_ context.Context, req create.ChatReq) (*create.ChatResp, error) {
	return &create.ChatResp{Name: req.Name}, nil
}

func main() {
	lambda.Start(handle)
}
