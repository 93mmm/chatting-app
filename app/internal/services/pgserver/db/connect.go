package db

import (
	"context"
	"log"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/helpers"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var conn *pgx.Conn

func CreateConnection() {
    var err error
	conn, err = pgx.Connect(context.Background(), helpers.GetDatabaseUrl())
	if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
	}
    log.Println("Connected to database")
}

func Exec(query string) (pgconn.CommandTag, error) {
    return conn.Exec(context.Background(), query)
}

func Close() {
    conn.Close(context.Background())
}
