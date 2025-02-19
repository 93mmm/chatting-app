package chat

import (
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/db"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
    var chat models.ChatAdd

    if err := chat.GetData(c); err != nil {
        helpers.WrapAndSendError(c, err)
        return 
    }

    id, err := db.CreateChat(chat)
    if err != nil {
        helpers.WrapAndSendError(c, err)
        return
    }

    err = db.AddParticipants(id, chat)
    if err != nil {
        helpers.WrapAndSendError(c, err)
        return
    }

    helpers.WrapAndSendId(c, id)
}
