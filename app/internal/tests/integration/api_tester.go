package integration

import (
	"testing"

	"github.com/93mmm/chatting-app/app/internal/tests/helpers"
	"github.com/93mmm/chatting-app/app/internal/tests/integration/chats"
	"github.com/93mmm/chatting-app/app/internal/tests/integration/users"
)

func TestApi(t *testing.T) {
    helpers.ClearDatabase()
    t.Cleanup(helpers.ClearDatabase)

    t.Run("TestUsers", users.Create)

    t.Run("TestChats", chats.Create)
    t.Run("TestChats", chats.Get)
}
