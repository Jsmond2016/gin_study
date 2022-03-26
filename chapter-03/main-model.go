package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Name     string `form:"name" json:"name" uri:"name" xml:"name"  binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	// 1.binding JSON
	// Example for binding JSON ({"user": "hanru", "password": "hanru123"})
	router.POST("/login-json", func(c *gin.Context) {
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

	// Form 表单绑定
	router.POST("/login-form", func(c *gin.Context) {
		var form Login
		//方法一：对于FORM数据直接使用Bind函数, 默认使用使用form格式解析,if c.Bind(&form) == nil
		// 根据请求头中 content-type 自动推断.
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.Name != "Tony" || form.Password != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "登录成功"})
	})

	// URI 绑定
	router.GET("/login/:name/:password", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": login.Name, "password": login.Password})
	})

	router.Run(":8080")
}
