package server

import (
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/server/api/chat"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/server/api/user"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
    {
        chatApi := router.Group("/api/chat")
        chatApi.POST("/create", chat.Add)
    }
    {
        userApi := router.Group("/api/user")
        userApi.POST("/reg", user.Reg)
    }
}
