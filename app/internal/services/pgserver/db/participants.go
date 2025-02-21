package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
)

func AddCreator(chatId int, chat models.ChatAdd) (error) {
    query := "INSERT INTO ChatParticipants (userId, chatId, role) VALUES ($1, $2, $3)"
    _, err := conn.Exec(context.Background(), query, chat.CreatorId, chatId, "c")
    return err
}

func AddParticipants(chatId int, chat models.ChatAdd) (error) {
    if !ExistsChat(chatId) {
        return errors.New("Chat doesn't exists")
    }
    for _, m := range chat.Members {
        query := "INSERT INTO ChatParticipants (userId, chatId) VALUES ($1, $2) ON CONFLICT DO NOTHING"
        _, err := conn.Exec(context.Background(), query, m, chatId)
        if err != nil {
            fmt.Println("[FUCKING ERROR]", err)
            fmt.Println(m, chatId)
        }
    }
    return nil
}
