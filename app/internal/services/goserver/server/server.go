package server

import (
	"fmt"
	"time"

	"github.com/93mmm/chatting-app/app/pkg/env"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func Run() {
    router := gin.Default()
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:8000"},
        AllowMethods:     []string{"GET", "PUT", "POST"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))
    setupRoutes(router)
    router.Run(fmt.Sprintf("0.0.0.0:%v", env.GoMainServer.Port))
}
