package server

import (
	"fmt"

	"github.com/93mmm/chatting-app/app/pkg/env"
	"github.com/gin-gonic/gin"
)

func Run() {
    router := gin.Default()
    setupRoutes(router)
    router.Run(fmt.Sprintf("0.0.0.0:%v", env.GoDatabaseServer.Port))
}
