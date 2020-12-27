package post

import "github.com/eyebrow-fish/orbit/chat"

type ChatReq struct {
	ChatId   int
	Body     string
	AuthorId int
}

type ChatResp struct {
	Message chat.Message
}
