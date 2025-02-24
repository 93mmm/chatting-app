package db

import (
	"context"
	"log"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
	se "github.com/93mmm/chatting-app/app/pkg/server_errors"
)

func CreateChat(chat models.ChatCreate) (int, error) {
    var chatId int
    query := "INSERT INTO Chats (creatorId, name, description) VALUES ($1, $2, $3) RETURNING id"

    if err := conn.QueryRow(context.Background(), query, chat.CreatorId, chat.Name, chat.Description).Scan(&chatId); err != nil {
        return 0, se.UndefinedError(err)
    }
    return chatId, nil
}

func AddChatCreator(chatId int, creatorId int) error {
    query := "INSERT INTO ChatParticipants (userId, chatId, role) VALUES ($1, $2, $3)"
    if _, err := conn.Exec(context.Background(), query, creatorId, chatId, string('c')); err != nil {
        return se.UndefinedError(err)
    }
    return nil
}

func AddChatMembers(chatId int, members []int) {
    for _, m := range members {
        query := "INSERT INTO ChatParticipants (userId, chatId, role) VALUES ($1, $2, $3)"
        if _, err := conn.Exec(context.Background(), query, m, chatId, string('m')); err != nil {
            log.Println(se.UndefinedError(err))
        }
    }
}
