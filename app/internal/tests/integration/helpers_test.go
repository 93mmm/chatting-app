package integration

import (
	"bytes"
	"io"
	"net/http"
)

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

