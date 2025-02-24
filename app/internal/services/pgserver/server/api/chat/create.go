package chat

import (
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/db"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
	se "github.com/93mmm/chatting-app/app/pkg/server_errors"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
    var chat models.ChatCreate

    if err := chat.GetData(c); err != nil {
        se.UndefinedError(err).Send(c)
        return
    }
    chatId, err := db.CreateChat(chat)
    if  err != nil {
        se.UndefinedError(err).Send(c)
        return
    }
    if err = db.AddChatCreator(chatId, chat.CreatorId); err != nil {
        se.UndefinedError(err).Send(c)
        return
    }
    db.AddChatMembers(chatId, chat.Members)
    helpers.WrapAndSendId(c, chatId)
}
