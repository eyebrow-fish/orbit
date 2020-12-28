package main

import (
	"database/sql"
	"github.com/eyebrow-fish/orbit/chat/post"
	"github.com/eyebrow-fish/orbit/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_handle_postNew(t *testing.T) {
	start := time.Now()
	ctx, err := testutil.DbCtx(t)
	assert.Nil(t, err)
	db := ctx.Value("db").(*sql.DB)
	_, err = db.Exec("create table chat(id serial primary key, name text not null)")
	assert.Nil(t, err)
	_, err = db.Exec(`
	create table message(
		id serial primary key, 
		chatId int not null,
		body text not null, 
		timestamp bigint not null,
		constraint fkChatId foreign key(chatId)
			references chat(id)
	)
	`)
	assert.Nil(t, err)
	_, err = db.Exec("insert into chat values(1, 'chat')")
	assert.Nil(t, err)
	t.Run("post new", func(t *testing.T) {
		res, err := handle(ctx, post.ChatReq{ChatId: 1, Body: "hello!"})
		assert.Nil(t, err)
		assert.LessOrEqual(t, start.UnixNano(), res.Message.Timestamp)
		assert.Equal(t, 1, res.Message.ChatId)
		assert.Equal(t, "hello!", res.Message.Body)
	})
}