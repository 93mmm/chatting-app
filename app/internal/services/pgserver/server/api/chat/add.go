package chat

import (
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/db"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
    var chatInfo models.Chat
    
    if err := chatInfo.GetData(c); err != nil {
        helpers.WrapAndSendError(c, err)
        return 
    }

    id, err := db.CreateChat(chatInfo)
    if err != nil {
        helpers.WrapAndSendError(c, err)
        return
    }

    helpers.WrapAndSendId(c, id)
}
