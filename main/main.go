package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
