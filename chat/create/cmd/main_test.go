package main

import (
	"github.com/eyebrow-fish/orbit/chat"
	"github.com/eyebrow-fish/orbit/chat/create"
	"github.com/eyebrow-fish/orbit/store"
	"github.com/eyebrow-fish/orbit/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_handle_createNew(t *testing.T) {
	dbCtx, err := testutil.DbCtx(t)
	assert.Nil(t, err)
	db := dbCtx.Value("db").(*store.Db)
	_, err = db.Exec("create table chat(id serial primary key, name text not null)")
	assert.Nil(t, err)
	t.Run("create new chat", func(t *testing.T) {
		resp, err := handle(dbCtx, create.ChatReq{Name: "new"})
		assert.Nil(t, err)
		assert.Equal(t, chat.Chat{Id: 1, Name: "new"}, resp.Chat)
	})
}

func Test_handle_createDuplicate(t *testing.T) {
	dbCtx, err := testutil.DbCtx(t)
	assert.Nil(t, err)
	db := dbCtx.Value("db").(*store.Db)
	_, err = db.Exec("create table chat(id serial primary key, name text not null)")
	t.Run("create duplicate", func(t *testing.T) {
		_, err = handle(dbCtx, create.ChatReq{Name: "dup"})
		assert.Nil(t, err)
		_, err = handle(dbCtx, create.ChatReq{Name: "dup"})
		assert.NotNil(t, err)
	})
}
