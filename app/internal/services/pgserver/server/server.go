package server

import (
	"fmt"

	"github.com/93mmm/chatting-app/app/internal/services/pgserver/server/api/chat"
	chat_edit "github.com/93mmm/chatting-app/app/internal/services/pgserver/server/api/chat/edit"
	chat_users "github.com/93mmm/chatting-app/app/internal/services/pgserver/server/api/chat/users"
	"github.com/93mmm/chatting-app/app/internal/services/pgserver/server/api/user"
	user_change "github.com/93mmm/chatting-app/app/internal/services/pgserver/server/api/user/change"
	"github.com/93mmm/chatting-app/app/pkg/env"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
    {
        chatApi := router.Group("/api/chat")
        
        chatApi.POST("/create", chat.Create)
        chatApi.GET("/get/:id", chat.GetById)
    }

    {
        chatApiUsers := router.Group("/api/chat/users")
        
        chatApiUsers.POST("/add", chat_users.Add)
        chatApiUsers.DELETE("/rm", chat_users.Remove)
        chatApiUsers.PUT("/promote", chat_users.Promote)
    }

    {
        chatApiEdit := router.Group("/api/chat/edit")

        chatApiEdit.PUT("/description", chat_edit.Description)
        chatApiEdit.PUT("/name", chat_edit.Description)
    }

    {
        userApi := router.Group("/api/user")

        userApi.POST("/reg", user.Reg) // completed
        userApi.GET("/get/:id", user.GetById) // completed
    }

    {
        userApiChange := router.Group("/api/user/change")

        userApiChange.PUT("/username", user_change.Username)
        userApiChange.PUT("/email", user_change.Email)
        userApiChange.PUT("/password", user_change.Password)
    }
}

func Run() {
    router := gin.Default()
    setupRoutes(router)
    router.Run(fmt.Sprintf("0.0.0.0:%v", env.GoDatabaseServer.Port))
}
