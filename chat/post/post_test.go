package post

import (
	"github.com/eyebrow-fish/orbit/store"
	"github.com/eyebrow-fish/orbit/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHandle_postNew(t *testing.T) {
	start := time.Now()
	ctx, _ := testutil.DbCtx(t)
	db := ctx.Value("db").(*store.Db)
	_, _ = db.Exec("create table Chat(Id serial primary key, Name text not null)")
	_, _ = db.Exec(`
	create table Message(
		Id serial primary key, 
		ChatId int not null,
		Body text not null, 
		Timestamp bigint not null,
		constraint FkChatId foreign key(ChatId)
			references Chat(Id)
	)
	`)
	_, err := db.Exec("insert into Chat values(1, 'chat')")
	assert.Nil(t, err)
	res, err := Handle(ctx, ChatReq{ChatId: 1, Body: "hello!"})
	assert.Nil(t, err)
	assert.LessOrEqual(t, start.UnixNano(), res.Message.Timestamp)
	assert.Equal(t, 1, res.Message.ChatId)
	assert.Equal(t, "hello!", res.Message.Body)
}

func TestHandle_noChatId(t *testing.T) {
	ctx, _ := testutil.DbCtx(t)
	db := ctx.Value("db").(*store.Db)
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
	_, err := Handle(ctx, ChatReq{ChatId: 1, Body: "failure!"})
	assert.NotNil(t, err)
}
