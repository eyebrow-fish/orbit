package main

import (
	"github.com/eyebrow-fish/orbit/chat/create"
	"github.com/eyebrow-fish/orbit/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_handle_createNew(t *testing.T) {
	dbCtx, err := testutil.DbCtx(t)
	assert.Nil(t, err)
	t.Run("create new chat", func(t *testing.T) {
		resp, err := handle(dbCtx, create.ChatReq{Name: "new"})
		assert.Nil(t, err)
		assert.Equal(t, "new", resp.Name)
	})
}
