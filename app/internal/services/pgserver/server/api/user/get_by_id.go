package user

import (
	"net/http"
	"strconv"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/db"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	se "github.com/93mmm/chatting-app/app/pkg/server_errors"
	"github.com/gin-gonic/gin"
)

func GetById(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        se.WrongValueError(c.Param("id")).Send(c)
        return
    }

    user, err := db.GetUserInfo(id)
    if err != nil {
        se.UndefinedError(err).Send(c)
        return
    }
    helpers.SendStruct(c, http.StatusOK, user)
}
