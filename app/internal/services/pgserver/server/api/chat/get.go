package chat

import (
	"net/http"
	"strconv"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
	"github.com/gin-gonic/gin"
)

// TODO: end this later
func Get(c *gin.Context) {
    var chatInfo models.ChatInfo
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        helpers.WrapAndSendError(c, err)
    }
    chatInfo.GetData(id)
    c.JSON(http.StatusOK, chatInfo)
}
