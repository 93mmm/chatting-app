package server

import (
	"log"
	"net/http"

	"github.com/93mmm/chatting-app/app/pkg/env"
	"github.com/93mmm/chatting-app/app/pkg/json"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
    data, err := json_rw.Read(env.CWD + "json/test.json")
    if err != nil {
        log.Println(err)
    }

    c.IndentedJSON(http.StatusOK, data)
}

func setupRoutes(router *gin.Engine) {
    router.GET("/test", test)
}
