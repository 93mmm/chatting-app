package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/server"
	"github.com/93mmm/chatting-app/app/pkg/db"
	"github.com/93mmm/chatting-app/app/pkg/env"
	"github.com/jackc/pgx/v5"
)

func main() {
    env.Init()

	conn, err := pgx.Connect(context.Background(), db.GetUrl())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
    log.Println("Connected to database")

    server.Run()
}
