package main

import (
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/server"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/db"
	"github.com/93mmm/chatting-app/app/pkg/env"
)

func main() {
    env.Init()
    db.CreateConnection()
	defer db.Close()

    server.Run()
}
