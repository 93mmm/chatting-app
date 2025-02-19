package integration

import (
	"testing"

	"github.com/93mmm/chatting-app/app/internal/tests/integration"
)

func TestChat(t *testing.T) {
    t.Run("Api testing", integration.TestApi)
}
