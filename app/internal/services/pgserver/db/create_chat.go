package db

import (
	"context"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
)

func CreateChat(chat models.Chat) (int, error) {
    tx, err := conn.Begin(context.Background())
    if err != nil {
        return 0, err
    }
    defer tx.Rollback(context.Background())

    var id int
    query := "INSERT INTO Chats (creatorId, name, description) VALUES ($1, $2, $3) RETURNING id"
    
    err = tx.QueryRow(context.Background(), query, chat.CreatorId, chat.Name, chat.Description).Scan(&id)
    if err != nil {
        return 0, err
    }

    err = tx.Commit(context.Background())
    if err != nil {
        return 0, err
    }
    return id, nil
}
