package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, err error) {
    e := struct{Err string `json:"error"`}{err.Error()}
    c.JSON(http.StatusBadRequest, e)
}

func WrapAndSendId(c *gin.Context, id int) {
    e := struct{Id int `json:"id"`}{id}
    c.JSON(http.StatusCreated, e)
}

func SendStruct(c *gin.Context, status int, data any) {
    c.JSON(status, data)
}
