package chat_edit

import (
	"errors"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	"github.com/gin-gonic/gin"
)

// TODO
func Description(c *gin.Context) {
    helpers.SendError(c, errors.New("Not implemented yet"))
}
