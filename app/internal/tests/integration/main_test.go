package integration

import (
	"testing"
)

func TestMain(t *testing.T) {
    clearDatabaseAfterTests()
    t.Cleanup(clearDatabaseAfterTests)
    t.Run("TestUser", testUser)
    t.Run("TestChats", testChatsCreate)
}
