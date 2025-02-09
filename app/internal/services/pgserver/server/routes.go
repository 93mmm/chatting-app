package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
    c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("Hello"))
}

func setupRoutes(router *gin.Engine) {
    router.GET("/hello", hello)
}
