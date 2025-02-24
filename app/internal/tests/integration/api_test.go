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

    tests := []string{
        "user/reg.json",
        "user/get_by_id.json",
        "chat/create.json",
    }
    for _, jsonTest := range tests {
        runTestFile := func(t *testing.T) {
            for _, link := range helpers.ReadJsonTests(jsonTest) {
                runTestRequest := func(t *testing.T) {
                    responce, err := helpers.MakeRequest(link.Sent)
                    helpers.AssertExpected(t, link.Expected, *responce, err)
                }

                t.Run(link.Name, runTestRequest)
            }
        }

        t.Run(jsonTest, runTestFile)
    }
}
