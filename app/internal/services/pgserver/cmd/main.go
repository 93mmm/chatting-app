package main

import (
	"log"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/db"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/server"
	"github.com/93mmm/chatting-app/app/pkg/env"
)

func main() {
    if err := env.Init(); err != nil {
        log.Fatal(err)
    }
    db.CreateConnection()
	defer db.Close()

    server.Run()
}
