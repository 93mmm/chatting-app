package integration

import (
	"log"
	"testing"

	"github.com/93mmm/chatting-app/app/internal/tests/helpers"
	"github.com/93mmm/chatting-app/app/pkg/env"
)

func TestUser(t *testing.T) {
    if err := env.Init(); err != nil {
        log.Fatal(err)
    }
    helpers.ClearDatabase()
    // t.Cleanup(helpers.ClearDatabase)

    tests := []string{
        "user/reg.json",
        // "chat/create.json",
    }
    for _, jsonTest := range tests {
        t.Run(jsonTest, func (t *testing.T) {
            for _, link := range helpers.ReadJsonTests(jsonTest) {
                responce, err := helpers.MakeRequest(link.Sent)
                helpers.AssertExpected(t, link.Expected, *responce, err)
            }
        })
    }
}
