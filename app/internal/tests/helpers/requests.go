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
    var resp Responce
    resp.Code = responce.StatusCode

    err = json.Unmarshal(body, &resp.Body)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

func PushTestUsers() {
    url := GetDatabaseServerUrl("/api/user/reg")
    for i := 0; i < 100; i++ {
        var request Request
        request.Url = url
        request.Method = "POST"
        json.Unmarshal([]byte(fmt.Sprintf(`{"email":"awesomeemail%v","pwdhash":"awesomehash"}`, i)), &request.Body)

        MakeRequest(request)
    }
}
