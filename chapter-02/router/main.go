package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "GET METHOD")
	})

	router.GET("/get-params/:name", func(c *gin.Context) {
		username := c.Param("name")
		c.String(http.StatusOK, username)
	})

	router.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "POST METHOD")
	})
	router.PUT("/put", func(c *gin.Context) {
		c.String(http.StatusOK, "PUT METHOD")
	})
	router.DELETE("/delete", func(c *gin.Context) {
		c.String(http.StatusOK, "DELETE METHOD")
	})
	router.PATCH("/patch", func(c *gin.Context) {
		c.String(http.StatusOK, "PATCH METHOD")
	})
	router.HEAD("/head", func(c *gin.Context) {
		c.String(http.StatusOK, "HEAD METHOD")
	})
	router.OPTIONS("/options", func(c *gin.Context) {
		c.String(http.StatusOK, "OPTIONS METHOD")
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
