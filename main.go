package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建gin路由器
	r := gin.Default()

	// 打印启动信息
	fmt.Println("正在启动gin服务器...")
	fmt.Println("服务器将在 http://localhost:8080 启动")

	// 添加一个简单的打招呼路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "你好！欢迎使用我们的gin程序！",
			"status":  "success",
		})
	})

	// 添加另一个打招呼路由
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
			"greeting": "欢迎来到gin世界！",
		})
	})

	// 启动服务器在8080端口
	r.Run(":8080")
}