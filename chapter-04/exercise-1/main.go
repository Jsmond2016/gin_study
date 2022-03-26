package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {
	router := gin.Default()

	// gin.H is a shortcut for map[string]interface{}
	router.GET("/json-gin", func(c *gin.Context) {
		// c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": nil})
	})

	router.GET("/json", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"userName"`
			Message string `json:"msg"`
			Number  int    `json:"num"`
		}
		msg.Name = "Tony"
		msg.Message = "理发"
		msg.Number = 100
		// 注意 msg.Name 变成了 "user" 字段
		// 以下方式都会输出 :   {"userName": "Tony", "msg": "理发", "num": 100}
		c.JSON(http.StatusOK, msg)
	})

	router.GET("/json-xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"user": "Tony", "message": "hey", "status": http.StatusOK})
	})

	router.GET("/json-yml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "Tony", "status": http.StatusOK})
	})

	router.GET("/protobuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// The specific definition of protobuf is written in the testdata/protoexample file.
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// Note that data becomes binary data in the response
		// Will output protoexample.Test protobuf serialized data
		c.ProtoBuf(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
