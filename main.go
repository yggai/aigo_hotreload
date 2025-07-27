package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建gin路由器
	r := gin.Default()

	// 打印启动信息
	fmt.Println("正在启动gin服务器...")
	fmt.Println("服务器将在 http://localhost:9000 启动")

	// 添加根路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "热更新测试成功！代码已自动重载111",
			"status":    "success",
			"timestamp": "2024-01-01",
		})
	})

	// 添加另一个打招呼路由
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "Hello, World!",
			"greeting": "欢迎来到gin世界！",
		})
	})

	// 启动服务器在9000端口
	r.Run(":9000")
}
