package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/eyebrow-fish/orbit/account/create"
)

func handle(_ context.Context, req create.AccountReq) (*create.AccountResp, error) {
	return &create.AccountResp{Name: req.Name}, nil
}

func main() {
	lambda.Start(handle)
}
