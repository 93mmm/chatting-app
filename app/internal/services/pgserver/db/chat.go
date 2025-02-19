package db

import (
	"context"
	"fmt"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
)

func CreateChat(data models.ChatAdd) (int, error) {
    tx, err := conn.Begin(context.Background())
    if err != nil {
        return 0, err
    }
    defer tx.Rollback(context.Background())

    var id int
    query := "INSERT INTO Chats (creatorId, name, description) VALUES ($1, $2, $3) RETURNING id"
    
    err = tx.QueryRow(context.Background(), query, data.CreatorId, data.Name, data.Description).Scan(&id)
    if err != nil {
        return 0, err
    }

    err = tx.Commit(context.Background())
    if err != nil {
        return 0, err
    }
    return id, nil
}

func EditChat(data models.EditChat) (error) {
    tx, err := conn.Begin(context.Background())
    if err != nil {
        return err
    }
    defer tx.Rollback(context.Background())

    var id int
    var updateName, updateDesc string
    if data.Name != nil {
        updateName = fmt.Sprintf("name=%v", *data.Name)
    }
    if data.Description != nil {
        updateDesc = fmt.Sprintf("description=%v", *data.Description)
    }

    query := fmt.Sprintf("UPDATE Chats SET %v, %v WHERE id = $1", updateName, updateDesc)
    
    err = tx.QueryRow(context.Background(), query, data.ChatId).Scan(&id)
    if err != nil {
        return err
    }

    err = tx.Commit(context.Background())
    if err != nil {
        return err
    }
    return nil
}

func ExistsChat(chat int) bool {
    conn.QueryRow(context.Background(), "SELECT 1 FROM Chats WHERE id=$1", chat).Scan(&chat)
    return chat == 1
}
