package get

import (
	"github.com/eyebrow-fish/orbit/chat/create"
	"github.com/eyebrow-fish/orbit/chat/post"
	"github.com/eyebrow-fish/orbit/store"
	"github.com/eyebrow-fish/orbit/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandle(t *testing.T) {
	dbCtx, _ := testutil.DbCtx(t)
	db := dbCtx.Value("db").(*store.Db)
	_, _ = db.Exec("create table chat(id serial primary key, name text not null)")
	_, _ = db.Exec(`
	create table message(
		id serial primary key, 
		chatId int not null,
		body text not null, 
		timestamp bigint not null,
		constraint fkChatId foreign key(chatId)
			references chat(id)
	)
	`)
	chatResp, _ := create.Handle(dbCtx, create.ChatReq{Name: "chat0"})
	_, _ = post.Handle(dbCtx, post.ChatReq{ChatId: chatResp.Chat.Id, Body: "post0"})
	_, _ = post.Handle(dbCtx, post.ChatReq{ChatId: chatResp.Chat.Id, Body: "post1"})
	_, _ = post.Handle(dbCtx, post.ChatReq{ChatId: chatResp.Chat.Id, Body: "post2"})
	resp, err := Handle(dbCtx, ChatReq{chatResp.Chat.Id})
	assert.Nil(t, err)
	assert.Equal(t, 3, len(resp.Messages))
	assert.Contains(t, "post0", resp.Messages[0].Body)
	assert.Contains(t, "post1", resp.Messages[1].Body)
	assert.Contains(t, "post2", resp.Messages[2].Body)
}
