package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/93mmm/chatting-app/app/pkg/json"
	"github.com/93mmm/chatting-app/app/pkg/env"
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func test(c *gin.Context) {
    data, err := json_handlers.ReadJSON(env.CWD + "json/test.json")
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Printf("data: %v\n", data)

    c.IndentedJSON(http.StatusOK, data)
}

func runGin() {
    router := gin.Default()

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:8000"},
        // AllowOrigins:     []string{fmt.Sprintf("http://%v:%v", env.)},
        AllowMethods:     []string{"GET", "PUT", "POST"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))
    router.GET("/test", test)
    router.Run(fmt.Sprintf("0.0.0.0:%v", env.GoMainServer.Port))
}

func main() {
    err := env.Init()
    if err != nil {
        log.Fatal(err)
    }

    runGin()
}
