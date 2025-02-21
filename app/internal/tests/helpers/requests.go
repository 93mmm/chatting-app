package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func MakeRequest(r Request) (*Responce, error) {
    client := http.Client{}
    jsonBody, err := json.Marshal(r.Body)
    if err != nil {
        return nil, err
    }
    request, err := http.NewRequest(r.Method, GetDatabaseServerUrl(r.Url), bytes.NewBuffer(jsonBody))
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

func PushTestUsers() {
    url := GetDatabaseServerUrl("/api/user/reg")
    for i := 0; i < 100; i++ {
        MakeRequest(
            Request{
                Url: url,
                Method: "POST",
                Body: fmt.Sprintf(`{"email":"awesomeemail%v","pwdhash":"awesomehash"}`, i),
            },
        )
    }
}
