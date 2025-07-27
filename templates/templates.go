package templates

// GoModTemplate go.mod文件模板
const GoModTemplate = `module %s

go 1.24

require (
	github.com/gin-gonic/gin v1.9.1
)
`

// MainGoTemplate main.go文件模板
const MainGoTemplate = `package main

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
	fmt.Println("服务器将在 http://localhost:8888 启动")

	// 添加根路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "热更新测试成功！代码已自动重载",
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

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"time":   "2024-01-01 00:00:00",
		})
	})

	// API路由组
	api := r.Group("/api/v1")
	{
		api.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"users": []string{"Alice", "Bob", "Charlie"},
			})
		})

		api.POST("/users", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"message": "User created successfully",
			})
		})
	}

	// 启动服务器在8888端口
	r.Run(":8888")
}
`

// AirTomlTemplate .air.toml文件模板
const AirTomlTemplate = `root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "go build -o ./tmp/main ."
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  include_file = []
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_root = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  time = false

[misc]
  clean_on_exit = false

[screen]
  clear_on_rebuild = false
  keep_scroll = true
`

// GitignoreTemplate .gitignore文件模板
const GitignoreTemplate = `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with go test -c
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# Go workspace file
go.work

# Air temporary files
tmp/

# IDE files
.vscode/
.idea/
*.swp
*.swo
*~

# OS generated files
.DS_Store
.DS_Store?
._*
.Spotlight-V100
.Trashes
ehthumbs.db
Thumbs.db

# Log files
*.log
build-errors.log
`

// ReadmeTemplate README.md文件模板
const ReadmeTemplate = `# %s

一个使用 aigo_hotreload 创建的 Go 热重载项目。

## 🚀 快速开始

### 环境要求
- Go 1.24 或更高版本

### 安装依赖
` + "```bash\n" + `go mod tidy
` + "```\n\n" + `### 启动开发服务器
` + "```bash\n" + `air
` + "```\n\n" + `### 访问应用
- 主页: http://localhost:8888
- 健康检查: http://localhost:8888/health
- API示例: http://localhost:8888/api/v1/users

## 📁 项目结构

` + "```\n" + `%s/
├── main.go              # 主程序文件
├── go.mod              # Go模块依赖
├── .air.toml           # Air热重载配置
├── .gitignore          # Git忽略文件
├── config/             # nginx配置文件目录
│   ├── your-domain.com # nginx配置文件
│   └── setup-nginx.sh  # nginx配置脚本
├── scripts/            # 脚本目录
│   └── apply-ssl.sh    # SSL证书申请脚本
└── README.md           # 项目说明
` + "```" + `

## 🛠️ 开发说明

- 修改代码后会自动重新编译和重启
- 默认端口: 8888
- 支持热重载，提高开发效率

## 🌐 域名部署

### 1. 配置nginx
` + "```bash\n" + `# 编辑nginx配置文件
vim config/your-domain.com

# 运行nginx配置脚本
chmod +x config/setup-nginx.sh
./config/setup-nginx.sh your-domain.com 8888
` + "```\n\n" + `### 2. 申请SSL证书
` + "```bash\n" + `# 安装certbot
sudo apt update
sudo apt install -y certbot python3-certbot-nginx

# 申请SSL证书
chmod +x scripts/apply-ssl.sh
./scripts/apply-ssl.sh your-domain.com
` + "```\n\n" + `### 3. 启动服务
` + "```bash\n" + `# 启动热重载服务
air
` + "```\n\n" + `## 📝 API接口

### GET /
返回欢迎信息

### GET /health
健康检查接口

### GET /api/v1/users
获取用户列表

### POST /api/v1/users
创建新用户

---

由 [aigo_hotreload](https://github.com/yggai/aigo_hotreload) 生成
`