package helpers

import (
	"fmt"

	"github.com/93mmm/chatting-app/app/pkg/env"
)

func GetDatabaseServerUrl(route string) string {
    if route[0] != '/' {
        route = "/" + route
    }
    return fmt.Sprintf("http://%v:%v%v", env.GoDatabaseServer.ContName, env.GoDatabaseServer.Port, route)
}
