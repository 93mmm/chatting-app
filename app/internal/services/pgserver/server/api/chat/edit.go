package chat

import (
	"net/http"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/db"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
	"github.com/gin-gonic/gin"
)

func Edit(c *gin.Context) {
    var data models.EditChat

    if err := data.GetData(c); err != nil {
        helpers.WrapAndSendError(c, err)
        return
    }

    if err := db.EditChat(data); err != nil {
        helpers.WrapAndSendError(c, err)
    }
    c.Status(http.StatusNoContent)

}
