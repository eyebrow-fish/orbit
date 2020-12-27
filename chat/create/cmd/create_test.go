package main

import (
	"database/sql"
	"github.com/eyebrow-fish/orbit/chat/create"
	"github.com/eyebrow-fish/orbit/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_handle_createNew(t *testing.T) {
	dbCtx, err := testutil.DbCtx(t)
	assert.Nil(t, err)
	db := dbCtx.Value("db").(*sql.DB)
	_, err = db.Exec("create table chat(id serial primary key, name text not null)")
	assert.Nil(t, err)
	t.Run("create new chat", func(t *testing.T) {
		resp, err := handle(dbCtx, create.ChatReq{Name: "new"})
		assert.Nil(t, err)
		assert.Equal(t, "new", resp.Name)
		rows, err := db.Query("select * from chat")
		assert.Nil(t, err)
		var id int
		var name string
		rows.Next()
		err = rows.Scan(&id, &name)
		assert.Nil(t, err)
		assert.Equal(t, 1, id)
		assert.Equal(t, "new", name)
	})
}
