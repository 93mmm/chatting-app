package chat_users

import (
	"errors"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	"github.com/gin-gonic/gin"
)

// TODO
func Add(c *gin.Context) {
    helpers.SendError(c, errors.New("Not implemented yet"))
}
