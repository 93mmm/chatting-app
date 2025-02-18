package db

import (
	"context"
	"errors"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/models"
)

func checkIfExists(regUsr models.RegisterUser) (bool, error) {
    var result int
    err := conn.QueryRow(context.Background(), "SELECT 1 FROM Users WHERE email = $1 OR username = $2", regUsr.Email, regUsr.Username).Scan(&result)
    return result == 1, err
}

func RegisterUser(regUsr models.RegisterUser) (int, error) {
    exists, err := checkIfExists(regUsr)
    if exists {
        return 0, errors.New("User already exists")
    }

    tx, err := conn.Begin(context.Background())
    if err != nil {
        return 0, err
    }
    defer tx.Rollback(context.Background())

    var id int
    query := "INSERT INTO Users (email, username, pwdHash) VALUES ($1, $2, $3) RETURNING id"
    
    err = tx.QueryRow(context.Background(), query, regUsr.Email, regUsr.Username, regUsr.PasswordHash).Scan(&id)
    if err != nil {
        return 0, err
    }

    return id, tx.Commit(context.Background())
}
