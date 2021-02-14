package create

import (
	"context"
	"fmt"
	cognitoIdp "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/google/uuid"
)

type AccountReq struct {
	Name     string
	Password string
}

type AccountResp struct {
	Name string
}

func Handle(ctx context.Context, req AccountReq) (*AccountResp, error) {
	client := ctx.Value("cognito").(cognitoIdp.Client)
	getUserResp, err := client.AdminGetUser(ctx, &cognitoIdp.AdminGetUserInput{Username: &req.Name})
	// An error is desired here because if there is no error,
	// then a user with this username already exists.
	if err == nil {
		return nil, fmt.Errorf("user already exists with username {%s}", *getUserResp.Username)
	}
	tempPassword := uuid.New().String()
	createUserResp, err := client.AdminCreateUser(
		ctx,
		&cognitoIdp.AdminCreateUserInput{
			Username:          &req.Name,
			TemporaryPassword: &tempPassword,
		},
	)
	if err != nil {
		return nil, err
	}
	_, err = client.AdminSetUserPassword(
		ctx,
		&cognitoIdp.AdminSetUserPasswordInput{
			Username:  createUserResp.User.Username,
			Password:  &req.Password,
			Permanent: true,
		},
	)
	if err != nil {
		return nil, err
	}
	return &AccountResp{*createUserResp.User.Username}, nil
}
