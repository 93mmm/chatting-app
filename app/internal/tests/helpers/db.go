package helpers

import (
	"log"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/db"
	"github.com/93mmm/chatting-app/app/pkg/env"
)

func ClearDatabase() {
    if err := env.Init(); err != nil {
        log.Fatal(err)
    }
    db.CreateConnection()
    defer db.Close()
    db.Exec("TRUNCATE Messages, ChatParticipants, Chats, Users RESTART IDENTITY CASCADE;")
}
