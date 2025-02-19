package helpers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func request(method string, r Request) (*Responce, error) {
    client := http.Client{}
    request, err := http.NewRequest(method, r.Url, bytes.NewBufferString(r.JsonBody))
    if err != nil {
        return nil, err
    }
    responce, err := client.Do(request)
    if err != nil {
        return nil, err
    }
    defer request.Body.Close()
    body, err := io.ReadAll(responce.Body)
    if err != nil {
        return nil, err
    }
    return &Responce{responce.StatusCode, string(body)}, nil
}

func PostRequest(r Request) (*Responce, error) {
    return request("POST", r)
}

func GetRequest(r Request) (*Responce, error) {
    return request("GET", r)
}

func PushTestUsers() {
    url := GetDatabaseServerUrl("/api/user/reg")
    for i := 0; i < 100; i++ {
        PostRequest(Request{
            url,
            fmt.Sprintf(`{"email":"awesomeemail%v","pwdhash":"awesomehash"}`, i),
        })
    }
}
