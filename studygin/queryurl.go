package main

import (
	"github.com/gin-gonic/gin"
	"log"
	_ "net/http"
)

//// 1. 获取url
//func main() {
//	// 初始化引擎
//	router := gin.Default()
//
//	// 假定一个 url 为 /welcome?firstname=Jane&lastname=Doe,
//	// 我们想要获取 firstname的内容
//	router.GET("/form_post", func(c *gin.Context) {
//		// 获取参数内容
//		// 获取的所有参数的类型都是 string
//		// 如果不存在，使用第二个当默认值
//		firstname := c.DefaultQuery("firstname", "Guest")
//		// 获取参数内容，没有则返回空字串
//		lastname := c.Query("lastname")
//
//		c.String(http.StatusOK, "hello %s %s", firstname, lastname)
//	})
//
//	router.Run(":8080")
//}

//// 2. 获取表单和body内容
//func main() {
//	router := gin.Default()
//
//	router.POST("/form_post", func(c *gin.Context) {
//		// 获取 post 过来的 message 内容
//		// 获取的参数类型是string
//		message := c.PostForm("message")
//		// 如果不存在，使用第二个当默认内容
//		nick := c.DefaultPostForm("nick", "anonymous")
//
//		c.JSON(200, gin.H{
//			"status":  "posted",
//			"message": message,
//			"nick":    nick,
//		})
//	})
//	router.Run(":8080")
//}

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		// 获取原始字节
		d, err := c.GetRawData()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println((string(d)))
		c.String(200, "0k")
	})
	router.Run(":8080")
}
