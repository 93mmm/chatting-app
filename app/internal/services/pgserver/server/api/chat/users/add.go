package chat

import (
	"fmt"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
	"github.com/gin-gonic/gin"
)

// TODO: end this later
func Add(c *gin.Context) {
    var user models.MoveUserInChat
    if err := user.GetData(c); err != nil {
        helpers.WrapAndSendError(c, err)
    }
    fmt.Println(user)
}
