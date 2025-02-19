package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
)

func AddParticipants(chatId int, chat models.ChatAdd) (error) {
    if !ExistsChat(chatId) {
        return errors.New("Chat doesn't exists")
    }
    // add creator ('c' role)
    query := "INSERT INTO ChatParticipants (userId, chatId, role) VALUES ($1, $2, $3)"
    _, err := conn.Exec(context.Background(), query, chat.CreatorId, chatId, 'c') 
    if err != nil {
        fmt.Println("[NOT CRITICAL ERROR]", err)
    }
    for _, m := range chat.Members {
        query := "INSERT INTO ChatParticipants (userId, chatId) VALUES ($1, $2)"
        _, err := conn.Exec(context.Background(), query, m, chatId)
        if err != nil {
            fmt.Println("[NOT CRITICAL ERROR]", err)
        }
    }
    return nil
}
