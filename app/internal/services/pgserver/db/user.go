package db

import (
	"context"

	se "github.com/93mmm/chatting-app/app/pkg/server_errors"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
)

func GetUserInfo(id int) (models.UserInfo, error) {
    var info models.UserInfo
    query := "SELECT username, email, regAt FROM Users WHERE id=$1"

    info.Id = id
    err := conn.QueryRow(context.Background(), query, id).Scan(&info.Username, &info.Email, &info.RegAt)
    if err != nil {
        return info, se.UserNotFoundError(info.Id)
    }
    return info, nil
}
