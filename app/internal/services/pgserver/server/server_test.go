package server

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/db"
	"github.com/93mmm/chatting-app/app/pkg/env"
	"github.com/stretchr/testify/assert"
)

type answer struct {
    JSONBody string
    Status   int
}

type requestResponce struct {
    Url              string
    JSONBody         string
    ExpectedStatus   int
    ExpectedJSONBody string
}

func makeJSONRequest(method string, rr *requestResponce) (*answer, error) {
    client := &http.Client{}
    request, err := http.NewRequest(method, rr.Url, bytes.NewBufferString(rr.JSONBody))
    if err != nil {
        return nil, err
    }
    resp, err := client.Do(request)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    data, err := io.ReadAll(resp.Body)
    return &answer{string(data), resp.StatusCode}, nil
}

func testUser(t *testing.T) {
    rrs := []requestResponce{
        {
            "http://go-db-server:8081/api/user/reg",
            `{"email":"awesomeemail","pwdhash":"awesomehash"}`,
            http.StatusCreated,
            `{"id":1}`,
        },
        {
            "http://go-db-server:8081/api/user/reg",
            `{"email":"awesomeemail1","pwdhash":"awesomehash"}`,
            http.StatusCreated,
            `{"id":2}`,
        },
        {
            "http://go-db-server:8081/api/user/reg",
            `{"email":"awesomeemail2","pwdhash":"awesomehash","username":"usrnm"}`,
            http.StatusCreated,
            `{"id":3}`,
        },
        {
            "http://go-db-server:8081/api/user/reg",
            `{"email":"awesomeemail2","pwdhash":"awesomehash","username":"usrnme"}`,
            http.StatusBadRequest,
            `{"error":"User already exists"}`,
        },
        {
            "http://go-db-server:8081/api/user/reg",
            `{"pwdhash":"awesomehash","username":"usrnme"}`,
            http.StatusBadRequest,
            `{"error":"Invalid Email"}`,
        },
        {
            "http://go-db-server:8081/api/user/reg",
            `{"email":"awesomeemail3","pwdhash":""}`,
            http.StatusBadRequest,
            `{"error":"Invalid PasswordHash"}`,
        },
        {
            "http://go-db-server:8081/api/user/reg",
            `{"email":"awesomeemail3","pwdhash":"hash"}`,
            http.StatusCreated,
            `{"id":4}`,
        },
    }
    for _, rr := range rrs {
        resp, err := makeJSONRequest("POST", &rr)
        if assert.NoError(t, err) {
            assert.Equal(t, rr.ExpectedJSONBody, resp.JSONBody, "Wrong JSON answer")
            assert.Equal(t, rr.ExpectedStatus, resp.Status, "Wrong status code")
        }
    }
}

func clearDatabaseAfterTests() {
    if err := env.Init(); err != nil {
        log.Fatal(err)
    }
    db.CreateConnection()
    defer db.Close()
    db.Exec("TRUNCATE Messages, ChatParticipants, Chats, Users RESTART IDENTITY CASCADE;")
}

func TestServer(t *testing.T) {
    testUser(t)

    clearDatabaseAfterTests()
}
