package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type wrappedError struct {
    ErrMsg  string  `json:"error"`
}

func WrapAndSendError(c *gin.Context, err error) {
    e := wrappedError{ErrMsg: err.Error()}
    c.JSON(http.StatusBadRequest, e)
}

type wrappedId struct {
    Id      int     `json:"id"`
}

func WrapAndSendId(c *gin.Context, id int) {
    e := wrappedId{Id: id}
    c.JSON(http.StatusCreated, e)
}
