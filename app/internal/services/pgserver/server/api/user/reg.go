package user

import (
	"context"
	"errors"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/db"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
	se "github.com/93mmm/chatting-app/app/pkg/server_errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func Reg(c *gin.Context) {
    var regUser models.RegUser
    
    if err := regUser.GetData(c); err != nil {
        se.UndefinedError(err).Send(c)
        return 
    }

    data, err := db.MakeTransaction(func(tx pgx.Tx) (any, error) {
        var id int
        query := `INSERT INTO Users (username, email, pwdHash) SELECT $1, $2, $3 WHERE NOT EXISTS (SELECT 1 FROM Users WHERE username = $1 OR email = $2) RETURNING ID;`
        err := tx.QueryRow(context.Background(), query, regUser.Username, regUser.Email, regUser.PasswordHash).Scan(&id)
        if errors.Is(err, pgx.ErrNoRows) {
            return 0, se.UserAlreadyExists(regUser.Username, regUser.Email)
        }
        return id, err
    })

    if err != nil {
        se.UndefinedError(err).Send(c)
        return
    }

    id, converted := data.(int)
    if !converted {
        se.UndefinedError(errors.New("Wrong data received from database")).Send(c)
        return
    }
    helpers.WrapAndSendId(c, id)
}
