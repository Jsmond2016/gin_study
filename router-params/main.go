package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/get", func(c *gin.Context) {
		name := c.DefaultQuery("name", "匿名者") //可设置默认值
		//nickname := c.Query("nickname") // 是 c.Request.URL.Query().Get("nickname") 的简写
		c.String(http.StatusOK, fmt.Sprintf("Hello %s ", name))
	})

	router.Run()
}
