package post

import (
	"context"
	"github.com/eyebrow-fish/orbit/chat"
	"github.com/eyebrow-fish/orbit/store"
	"time"
)

type ChatReq struct {
	ChatId   int
	Body     string
}

type ChatResp struct {
	Message chat.Message
}

func Handle(ctx context.Context, req ChatReq) (*ChatResp, error) {
	postTime := time.Now().UnixNano()
	db := ctx.Value("db").(*store.Db)
	err := db.ExecUnique(
		`
		insert into Message(ChatId, Body, Timestamp) 
		select $1, $2, $3
		where not exists (
			select 1 from message where ChatId = $1 and Body = $2 and Timestamp = $3
		)
		and exists (
			select 1 from Chat where Id = $1
		)
		`,
		req.ChatId,
		req.Body,
		postTime,
	)
	if err != nil {
		return nil, err
	}
	msg, err := db.QueryUnique(
		chat.Message{},
		"select * from Message where ChatId = $1 and Timestamp = $2",
		req.ChatId,
		postTime,
	)
	return &ChatResp{Message: msg.(chat.Message)}, err
}
