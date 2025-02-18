package db

import (
	"context"
	"fmt"
	"log"

	"github.com/93mmm/chatting-app/app/pkg/env"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var conn *pgx.Conn

func getUrl() string {
    return fmt.Sprintf("postgres://%v:%v@%v:5432/%v",
                       env.PostgresDatabase.User,
                       env.PostgresDatabase.Password,
                       env.PostgresDatabase.ContName,
                       env.PostgresDatabase.DbName,
                       )
}

func CreateConnection() {
    var err error
	conn, err = pgx.Connect(context.Background(), getUrl())
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
