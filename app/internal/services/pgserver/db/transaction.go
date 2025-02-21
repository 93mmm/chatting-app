package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)


func MakeTransaction(fn func(transaction pgx.Tx) (any, error)) (any, error) {
    tx, err := conn.Begin(context.Background())
    if err != nil {
        return nil, err
    }
    defer tx.Rollback(context.Background())

    result, err := fn(tx)
    if err != nil {
        fmt.Println(err.Error())
        return nil, err
    }

    err = tx.Commit(context.Background())
    return result, err
}
