package helpers

import (
	"net/http"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/gin-gonic/gin"
)

func parseError(err error) error {
    if err != nil {
        var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) {
            log.Println(err)
            return errors.New("Internal database error")
        }
        return err
    }
    return nil
}

type wrappedError struct {
    ErrMsg  string  `json:"error"`
}

func WrapAndSendError(c *gin.Context, err error) {
    err = parseError(err)
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

