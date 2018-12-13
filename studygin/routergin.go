package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 1.路由
//func main() {
// 初始化引擎
//router := gin.Default()

// 注册的方法
//router.GET("/someGET",geting)
//router.POST("/somePOST",posting)
//router.PUT("/somePUT",puting)
//router.DELETE("/someDELETE",deleting)
//router.PATCH("/somePATCH",patching)
//router.HEAD("/someHEAD",heading)
//router.OPTIONS("/someOPTIONS",optionsing)

// 默认绑定 :8080
//router.run()
//}

//// 2.动态路由
//func main() {
//	router := gin.Default()
//
//	// 注册一个动态路由
//	// 可以匹配 /user/joy 不能匹配/user 和 /user/
//	router.Get("/user/:name", func(c *gin.Context) {
//		// 使用 c.Param(key) 获取 url 参数
//		name := c.Param("name")
//		c.String(http.StatusOK, "hello %s", name)
//	})
//
//	// 注册一个高级动态路由
//	// 可以匹配到 /user/john/ 和  /user/john/send
//	// 如果没有任何路由匹配到 /user/john，那么他就会重定向到  /user/john/ 从而被该匹配方法匹配到
//	router.GET("/Users/:name/*action", func(c *gin.Context) {
//		name := c.Param("name")
//		action := c.Param("action")
//		message := name + "is" + action
//		c.String(http.StatusOK, message)
//	})
//
//	router.Run(":8080")
//}

// 3.路由组
//func main(){
//router := gin.Default()

// 定义一个前缀，比如可以匹配到 /v1/fath
//v1 := router.Group("/v1")
//{
//v1.POST("/fath",fathPoint)
//v1.POST("/read",readPoint)
//}

// 其实也不用用花括号括起来，类似下面
//v2 := router.Group("/v2")
//v2.POST("/fath",fathpoint)
//v2.POST("/read",readpoint)

//router.Run(":8080")
//}
