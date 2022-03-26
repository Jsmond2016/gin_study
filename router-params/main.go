package main

import (
	"fmt"
	"log"
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

	//form
	router.POST("/form", func(c *gin.Context) {
		// type1 := c.DefaultPostForm("type", "alert") //可设置默认值
		username := c.PostForm("username")
		password := c.PostForm("password")

		//hobbys := c.PostFormMap("hobby")
		//hobbys := c.QueryArray("hobby")
		hobbys := c.PostFormArray("hobby")

		c.String(http.StatusOK, fmt.Sprintf("username is %s, password is %s,hobby is %v", username, password, hobbys))

	})

	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, file.Filename)

		/*
			也可以直接使用io操作，拷贝文件数据。
			out, err := os.Create(filename)
			defer out.Close()
			_, err = io.Copy(out, file)
		*/

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	//router.Static("/", "./public")
	router.POST("/uploads", func(c *gin.Context) {

		// Multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		files := form.File["files"]

		for _, file := range files {
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}

		c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files ", len(files)))
	})

	router.Run()
}
