package users

import (
	"net/http"
	"testing"

	"github.com/93mmm/chatting-app/app/internal/tests/helpers"
	"github.com/stretchr/testify/assert"
)

func assertExpected(t *testing.T, exp helpers.Responce, res helpers.Responce, err error) {
    if assert.NoError(t, err) {
        assert.Equal(t, exp.Code, res.Code, "Wrong status code")
        assert.Equal(t, exp.JsonBody, res.JsonBody, "Wrong JSON answer")
    }
}

func makeLinks() []helpers.ExpectedResponceToRequest {
    url := helpers.GetDatabaseServerUrl("/api/user/reg")
    return []helpers.ExpectedResponceToRequest{
        {
            Sent:       helpers.Request {Url: url, JsonBody: `{"email":"awesomeemail","pwdhash":"awesomehash"}`},
            Expected:   helpers.Responce{Code: http.StatusCreated, JsonBody: `{"id":1}`},
        },
        {
            Sent:       helpers.Request {Url: url, JsonBody: `{"email":"awesomeemail1","pwdhash":"awesomehash"}`},
            Expected:   helpers.Responce{Code: http.StatusCreated, JsonBody: `{"id":2}`},
        },
        {
            Sent:       helpers.Request {Url: url, JsonBody: `{"email":"awesomeemail2","pwdhash":"awesomehash","username":"usrnm"}`},
            Expected:   helpers.Responce{Code: http.StatusCreated, JsonBody: `{"id":3}`},
        },
        {
            Sent:       helpers.Request {Url: url, JsonBody: `{"email":"awesomeemail2","pwdhash":"awesomehash","username":"usrnme"}`},
            Expected:   helpers.Responce{Code: http.StatusBadRequest, JsonBody: `{"error":"User already exists"}`},
        },
        {
            Sent:       helpers.Request {Url: url, JsonBody: `{"pwdhash":"awesomehash","username":"usrnme"}`},
            Expected:   helpers.Responce{Code: http.StatusBadRequest, JsonBody: `{"error":"Invalid Email"}`},
        },
        {
            Sent:       helpers.Request {Url: url, JsonBody: `{"email":"awesomeemail3","pwdhash":""}`},
            Expected:   helpers.Responce{Code: http.StatusBadRequest, JsonBody: `{"error":"Invalid PasswordHash"}`},
        },
        {
            Sent:       helpers.Request {Url: url, JsonBody: `{"email":"awesomeemail3","pwdhash":"hash"}`},
            Expected:   helpers.Responce{Code: http.StatusCreated, JsonBody: `{"id":4}`},
        },
    }
}

func Create(t *testing.T) {
    for _, link := range makeLinks() {
        responce, err := helpers.PostRequest(link.Sent)
        helpers.AssertExpected(t, link.Expected, *responce, err)
    }
}
