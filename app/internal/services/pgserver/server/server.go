package server

import (
	"fmt"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/server/api/chat"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/server/api/user"
	"github.com/93mmm/chatting-app/app/pkg/env"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
    {
        chatApi := router.Group("/api/chat")
        chatApi.POST("/create", chat.Add)
        chatApi.PUT("/edit", chat.Edit)
    }
    {
        userApi := router.Group("/api/user")
        userApi.POST("/reg", user.Reg)
    }
}

func Run() {
    router := gin.Default()
    setupRoutes(router)
    router.Run(fmt.Sprintf("0.0.0.0:%v", env.GoDatabaseServer.Port))
}
