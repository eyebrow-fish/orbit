package create

import "github.com/eyebrow-fish/orbit/chat"

type ChatReq struct {
	Name string
}

type ChatResp struct {
	Chat chat.Chat
}
