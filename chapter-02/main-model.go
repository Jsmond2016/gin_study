package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Name     string `form:"name" json:"name" uri:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	//1.binding JSON
	// Example for binding JSON ({"user": "hanru", "password": "hanru123"})
	router.POST("/login", func(c *gin.Context) {
		var json Login
		//其实就是将request中的Body中的数据按照JSON格式解析到json变量中
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.Name != "Tony" || json.Password != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "登录成功"})
	})

	router.Run(":8080")
}
