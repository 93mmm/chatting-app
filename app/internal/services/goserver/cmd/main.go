package main

import (
	"log"

	"github.com/93mmm/chatting-app/app/internal/services/goserver/server"
	"github.com/93mmm/chatting-app/app/pkg/env"
)


func main() {
    err := env.Init()
    if err != nil {
        log.Fatal(err)
    }

    server.Run()
}
