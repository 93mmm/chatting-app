package main

import (
	"log"
	"net/http"
	"fmt"

	"github.com/93mmm/chatting-app/app/pkg/env"
)

func main() {
    err := env.Init()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(env.GoDatabaseServer)
    fmt.Println(env.PostgresDatabase)
    fmt.Println(env.GoDatabaseServer)
    fmt.Println(env.CWD)

    resp, err := http.Get(fmt.Sprintf("http://%v:%v/test",
                                      env.GoMainServer.ContName,
                                      env.GoMainServer.Port))
    if err != nil {
        log.Fatal(err)
    }
    log.Println(resp)
}
