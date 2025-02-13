package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/93mmm/chatting-app/app/pkg/env"
	"github.com/93mmm/chatting-app/app/pkg/json"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
    fmt.Println(env.CWD)
    data, err := json_rw.Read(env.CWD + "json/test.json")
    if err != nil {
        log.Println(err)
    }
    fmt.Printf("data: %v\n", data)

    c.IndentedJSON(http.StatusOK, data)
}

func setupRoutes(router *gin.Engine) {
    router.GET("/test", test)
}
