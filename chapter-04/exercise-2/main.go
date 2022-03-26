package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//--------- 单个 index 模板 -----------------
	// router.LoadHTMLGlob("templates")
	// router.GET("/index", func(c *gin.Context) {
	// 	//根据完整文件名渲染模板，并传递参数
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "Hello, World!",
	// 	})
	// })

	// -------- 多个 index 模板 -----------------
	// 多个 index 模板，使用 * 匹配多个层级，index 同名也没关系
	router.LoadHTMLGlob("templates/**/*")

	router.GET("/posts/index", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Hello, Posts!",
		})
	})

	router.GET("/user/index", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "user/index.tmpl", gin.H{
			"title": "Hello, User!",
		})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
