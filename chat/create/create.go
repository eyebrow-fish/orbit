package create

import (
	"context"
	"github.com/eyebrow-fish/orbit/chat"
	"github.com/eyebrow-fish/orbit/store"
)

type ChatReq struct {
	Name string
}

type ChatResp struct {
	Chat chat.Chat
}

func Handle(ctx context.Context, req ChatReq) (*ChatResp, error) {
	db := ctx.Value("db").(*store.Db)
	err := db.ExecUnique("insert into Chat(Name) values($1)", req.Name)
	if err != nil {
		return nil, err
	}
	resp, err := db.QueryUnique(chat.Chat{}, "select * from Chat where Name = $1", req.Name)
	return &ChatResp{resp.(chat.Chat)}, err
}
