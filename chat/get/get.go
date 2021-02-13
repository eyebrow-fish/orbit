package get

import (
	"context"
	"github.com/eyebrow-fish/orbit/chat"
	"github.com/eyebrow-fish/orbit/store"
)

type ChatReq struct {
	ChatId int
}

type ChatResp struct {
	Messages []chat.Message
}

func Handle(ctx context.Context, req ChatReq) (*ChatResp, error) {
	db := ctx.Value("db").(*store.Db)
	resp, err := db.QueryMany(
		chat.Message{},
		`select * from Message where ChatId = $1`,
		req.ChatId,
	)
	if err != nil {
		return nil, err
	}
	var messages []chat.Message
	for _, m := range resp {
		messages = append(messages, m.(chat.Message))
	}
	return &ChatResp{messages}, nil
}
