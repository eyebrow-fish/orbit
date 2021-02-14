package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	cognitoIdp "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/eyebrow-fish/orbit/account/create"
	"github.com/eyebrow-fish/orbit/store"
)

func main() {
	dbCtx, err := store.DbCtx()
	cognitoClient := cognitoIdp.New(cognitoIdp.Options{})
	cognitoCtx := context.WithValue(dbCtx, "cognito", cognitoClient)
	if err != nil {
		panic(err)
	}
	lambda.StartWithContext(cognitoCtx, create.Handle)
}
