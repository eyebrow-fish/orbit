package create

import (
	"github.com/eyebrow-fish/orbit/chat"
	"github.com/eyebrow-fish/orbit/store"
	"github.com/eyebrow-fish/orbit/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandle_createNew(t *testing.T) {
	dbCtx, _ := testutil.DbCtx(t)
	db := dbCtx.Value("db").(*store.Db)
	_, _ = db.Exec("create table chat(id serial primary key, name text not null)")
	resp, err := Handle(dbCtx, ChatReq{"new"})
	assert.Nil(t, err)
	assert.Equal(t, chat.Chat{Id: 1, Name: "new"}, resp.Chat)
}

func TestHandle_createDuplicate(t *testing.T) {
	dbCtx, _ := testutil.DbCtx(t)
	db := dbCtx.Value("db").(*store.Db)
	_, _ = db.Exec("create table chat(id serial primary key, name text not null)")
	_, err := Handle(dbCtx, ChatReq{"dup"})
	assert.Nil(t, err)
	_, err = Handle(dbCtx, ChatReq{"dup"})
	assert.Nil(t, err)
}
