package integration

import (
	"log"
	"testing"

	"github.com/93mmm/chatting-app/app/internal/tests/helpers"
	"github.com/93mmm/chatting-app/app/pkg/env"
)

func TestApi(t *testing.T) {
    if err := env.Init(); err != nil {
        log.Fatal(err)
    }
    helpers.ClearDatabase()

    tests := []struct{ file string; fields []string }{
        { "user/reg.json", []string{"id"} },
        { "user/get_by_id.json", []string{"id", "username", "email"} },
        { "chat/create.json", []string{"id"} },
        { "chat/get_by_id.json", []string{"chatId", "creatorId", "name", "description"} },
    }

    for _, test := range tests {
        runTestFile := func(t *testing.T) {
            for _, link := range helpers.ReadJsonTests(test.file) {
                runTestRequest := func(t *testing.T) {
                    responce, err := helpers.MakeRequest(link.Sent)
                    helpers.AssertExpected(t, link.Expected, *responce, err, test.fields)
                }

                t.Run(link.Name, runTestRequest)
            }
        }

        t.Run(test.file, runTestFile)
    }
}
