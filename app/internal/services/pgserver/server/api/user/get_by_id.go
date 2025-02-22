package user

import (
	"net/http"
	"strconv"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/db"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	"github.com/gin-gonic/gin"
)


func GetById(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        helpers.SendError(c, err)
        return
    }

    user, err := db.GetUserInfo(id)
    if err != nil {
        helpers.SendError(c, err)
        return
    }
    helpers.SendStruct(c, http.StatusOK, user)
}
