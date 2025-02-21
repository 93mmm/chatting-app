package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
)

func GetUserInfo(id int) (models.UserInfo, error) {
    var info models.UserInfo
    query := "SELECT username, email, regAt FROM Users WHERE id=$1"

    info.Id = id
    err := conn.QueryRow(context.Background(), query, id).Scan(&info.Username, &info.Email, &info.RegAt)
    if err != nil {
        fmt.Println(err)
        return info, errors.New("Not found specified user")
    }
    return info, nil
}
