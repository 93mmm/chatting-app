package server

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/db"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	"github.com/93mmm/chatting-app/app/pkg/env"
	"github.com/stretchr/testify/assert"
)

type answer struct {
    JSONBody string
    Status   int
}

type requestResponce struct {
    JSONBody         string
    ExpectedStatus   int
    ExpectedJSONBody string
}

func makeGetRequest(url string) (*answer, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    data, err := io.ReadAll(resp.Body)
    return &answer{string(data), resp.StatusCode}, nil
}

func makeJSONRequest(method, url string, rr *requestResponce) (*answer, error) {
    client := &http.Client{}
    request, err := http.NewRequest(method, url, bytes.NewBufferString(rr.JSONBody))
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

func testChatsCreate(t *testing.T) {
    if err := env.Init(); err != nil {
        log.Fatal(err)
    }
    url := helpers.GetDatabaseServerUrl("/api/chat/create") 
    rrs := []requestResponce{
        {
            `{"creatorId":1,"members":[1, 2, 3],"name":"Awesome Chat","description":"Awesome Chat Description"}`,
            http.StatusCreated,
            `{"id":1}`,
        },
        {
            `{"creatorId":1,"members":[1, 2, 3],"name":"Awesome Chat","description":"Awesome Chat Description"}`,
            http.StatusCreated,
            `{"id":2}`,
        },
        {
            `{"creatorId":100000,"members":[100000],"name":"Awesome Chat","description":"Awesome Chat Description"}`,
            http.StatusBadRequest,
            `{"error":"Internal database error"}`,
        },
    }

    for _, rr := range rrs {
        resp, err := makeJSONRequest("POST", url, &rr)
        if assert.NoError(t, err) {
            assert.Equal(t, rr.ExpectedJSONBody, resp.JSONBody, "Wrong JSON answer")
            assert.Equal(t, rr.ExpectedStatus, resp.Status, "Wrong status code")
        }
    }
}

// func testChatsGet(t *testing.T) {
//     if err := env.Init(); err != nil {
//         log.Fatal(err)
//     }
//     url := helpers.GetDatabaseServerUrl("/api/chat/get")
//     rrs := []requestResponce{
//         {
//             "",
//             http.StatusOK,
//             `{"id":1}`,
//         },
//     }
//     for _, rr := range rrs {
//         makeGetRequest()
//     }
// }
//
func testUser(t *testing.T) {
    if err := env.Init(); err != nil {
        log.Fatal(err)
    }
    rrs := []requestResponce{
        {
            `{"email":"awesomeemail","pwdhash":"awesomehash"}`,
            http.StatusCreated,
            `{"id":1}`,
        },
        {
            `{"email":"awesomeemail1","pwdhash":"awesomehash"}`,
            http.StatusCreated,
            `{"id":2}`,
        },
        {
            `{"email":"awesomeemail2","pwdhash":"awesomehash","username":"usrnm"}`,
            http.StatusCreated,
            `{"id":3}`,
        },
        {
            `{"email":"awesomeemail2","pwdhash":"awesomehash","username":"usrnme"}`,
            http.StatusBadRequest,
            `{"error":"User already exists"}`,
        },
        {
            `{"pwdhash":"awesomehash","username":"usrnme"}`,
            http.StatusBadRequest,
            `{"error":"Invalid Email"}`,
        },
        {
            `{"email":"awesomeemail3","pwdhash":""}`,
            http.StatusBadRequest,
            `{"error":"Invalid PasswordHash"}`,
        },
        {
            `{"email":"awesomeemail3","pwdhash":"hash"}`,
            http.StatusCreated,
            `{"id":4}`,
        },
    }
    url := helpers.GetDatabaseServerUrl("/api/user/reg")
    for _, rr := range rrs {
        resp, err := makeJSONRequest("POST", url, &rr)
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
    clearDatabaseAfterTests()
    t.Cleanup(clearDatabaseAfterTests)
    t.Run("TestUser", testUser)
    t.Run("TestChats", testChatsCreate)
    // t.Run("TestChats", testChatsGet)
    // t.Run("TestChatUsers", testChatsEdit)
}
