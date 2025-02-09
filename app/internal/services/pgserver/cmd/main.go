package main

import (
	"context"
	"fmt"
	"os"

	"github.com/93mmm/chatting-app/app/pkg/env"
	"github.com/jackc/pgx/v5"
)

func main() {
    env.Init()

    url := fmt.Sprintf("postgres://%v:%v@%v:5432/%v",
                       env.PostgresDatabase.User,
                       env.PostgresDatabase.Password,
                       env.PostgresDatabase.ContName,
                       env.PostgresDatabase.DbName,
                       )
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
    fmt.Println("Connected to database")
}
