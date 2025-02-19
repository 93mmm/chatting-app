package helpers

import (
	"fmt"

	"github.com/93mmm/chatting-app/app/pkg/env"
)

func GetDatabaseUrl() string {
    return fmt.Sprintf("postgres://%v:%v@%v:5432/%v",
                       env.PostgresDatabase.User,
                       env.PostgresDatabase.Password,
                       env.PostgresDatabase.ContName,
                       env.PostgresDatabase.DbName,
                       )
}

func GetDatabaseServerUrl(route string) string {
    if route[0] != '/' {
        route = "/" + route
    }
    return fmt.Sprintf("http://%v:%v%v", env.GoDatabaseServer.ContName, env.GoDatabaseServer.Port, route)
}
