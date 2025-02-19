package chats

import (
	"net/http"
	"testing"

	"github.com/93mmm/chatting-app/app/internal/tests/helpers"
)

func makeLinks() []helpers.ExpectedResponceToRequest {
    url := helpers.GetDatabaseServerUrl("/api/chat/create")
    return []helpers.ExpectedResponceToRequest{
        {
            Sent:       helpers.Request {Url: url, JsonBody: `{"creatorId":1,"members":[1, 2, 3],"name":"Awesome Chat","description":"Awesome Chat Description"}`},
            Expected:   helpers.Responce{Code: http.StatusCreated, JsonBody: `{"id":1}`},
        },
        {
            Sent:       helpers.Request {Url: url, JsonBody: `{"creatorId":1,"members":[1, 2, 3],"name":"Awesome Chat","description":"Awesome Chat Description"}`},
            Expected:   helpers.Responce{Code: http.StatusCreated, JsonBody: `{"id":2}`},
        },
        {
            Sent:       helpers.Request {Url: url, JsonBody: `{"creatorId":100000,"members":[100000],"name":"Awesome Chat","description":"Awesome Chat Description"}`},
            Expected:   helpers.Responce{Code: http.StatusBadRequest, JsonBody: `{"error":"Internal database error"}`},
        },
    }
}

func Create(t *testing.T) {
    // helpers.PushTestUsers()
    for _, link := range makeLinks() {
        responce, err := helpers.PostRequest(link.Sent)
        helpers.AssertExpected(t, link.Expected, *responce, err)
    }
}
