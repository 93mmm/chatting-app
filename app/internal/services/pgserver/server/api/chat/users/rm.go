package chat

import (
	"fmt"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
	"github.com/gin-gonic/gin"
)

// TODO: end this later
func Remove(c *gin.Context) {
    var user models.MoveUserInChat
    if err := user.GetData(c); err != nil {
        helpers.SendError(c, err)
    }
    fmt.Println(user)
}
