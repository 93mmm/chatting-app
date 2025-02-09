package db

import (
	"fmt"

	"github.com/93mmm/chatting-app/app/pkg/env"
)


func GetUrl() string {
    return fmt.Sprintf("postgres://%v:%v@%v:5432/%v",
                       env.PostgresDatabase.User,
                       env.PostgresDatabase.Password,
                       env.PostgresDatabase.ContName,
                       env.PostgresDatabase.DbName,
                       )
}

