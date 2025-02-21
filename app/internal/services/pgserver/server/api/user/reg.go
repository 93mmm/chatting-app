package user

import (
	"context"
	"errors"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/db"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func Reg(c *gin.Context) {
    var regUser models.RegUser
    
    if err := regUser.GetData(c); err != nil {
        helpers.SendError(c, err)
        return 
    }

    data, err := db.MakeTransaction(func(tx pgx.Tx) (any, error) {
        var id int
        query := `INSERT INTO Users (username, email, pwdHash) SELECT $1, $2, $3 WHERE NOT EXISTS (SELECT 1 FROM Users WHERE username = $1 OR email = $2) RETURNING ID;`
        err := tx.QueryRow(context.Background(), query, regUser.Username, regUser.Email, regUser.PasswordHash).Scan(&id)
        if err != nil {
            return 0, errors.New("User already exists")
        }
        return id, nil
    })

    if err != nil {
        helpers.SendError(c, err)
        return 
    }

    id, converted := data.(int)
    if !converted {
        helpers.SendError(c, errors.New("Wrong data received from database"))
        return
    }
    helpers.WrapAndSendId(c, id)
}
