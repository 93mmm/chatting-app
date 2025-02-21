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
        helpers.SendError(c, err)
        return 
    }

    id, err := db.CreateChat(chat)
    if err != nil {
        helpers.SendError(c, err)
        return
    }

    if err = db.AddCreator(id, chat); err != nil {
        helpers.SendError(c, err)
        return
    }
    
    if err = db.AddParticipants(id, chat); err != nil {
        helpers.SendError(c, err)
        return
    }

    helpers.WrapAndSendId(c, id)
}
