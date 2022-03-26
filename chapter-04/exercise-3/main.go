package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 下面测试静态文件服务

	// router.StaticFS("/files", http.Dir("../"))

	// Static 提供给定文件系统根目录中的文件。
	router.Static("/assets", "./assets") // 请求方式：http://localhost:8080/assets/img.png
	// 显示当前文件夹下的所有文件/或者指定文件
	router.StaticFS("/show-dir", http.Dir("."))
	router.StaticFile("/image", "./images/img.png")

	router.GET("/redirect", func(c *gin.Context) {
		//支持内部和外部的重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	router.Run(":8080")
}
