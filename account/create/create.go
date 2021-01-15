package create

import "context"

type AccountReq struct {
	Name string
}

type AccountResp struct {
	Name string
}

func Handle(_ context.Context, req AccountReq) (*AccountResp, error) {
	return &AccountResp{Name: req.Name}, nil
}
