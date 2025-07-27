# Go 热重载开发模板 (aigo_hotreload)

> 🚀 源滚滚AI编程开发的Go语言热重载模板项目，专为云服务器开发环境设计，支持本地和域名热更新，大幅提升开发效率！

[中文](README.md) | [English](README_EN.md)

---

### 📋 项目简介

本项目是一个完整的Go语言热重载开发模板，集成了以下特性：
- ✅ 本地热重载开发
- ✅ 域名反向代理配置
- ✅ HTTPS SSL证书自动配置
- ✅ Nginx反向代理支持
- ✅ 生产环境部署指南

### 🛠️ 技术栈

- **后端框架**: Gin (Go)
- **热重载工具**: Air
- **反向代理**: Nginx
- **SSL证书**: Let's Encrypt (免费)
- **部署环境**: Linux云服务器

### 📁 项目结构

```
aigo_hotreload/
├── main.go              # CLI工具主程序
├── server.go            # 示例服务器文件
├── go.mod              # Go模块依赖
├── go.sum              # 依赖校验文件
├── .air.toml           # Air热重载配置
├── config/             # 配置文件目录
│   └── testapi.zhangdapeng.com  # Nginx站点配置
├── docs/               # 文档目录
├── README.md           # 项目说明文档（中文）
└── README_EN.md        # 项目说明文档（英文）
```

### 🚀 快速开始

#### 环境要求
- Go 1.24 或更高版本

#### 1. 克隆项目
```bash
git clone https://github.com/yggai/aigo_hotreload.git
cd aigo_hotreload
```

#### 2. 安装依赖
```bash
go mod tidy
```

#### 3. 本地开发（热重载）
```bash
# 安装Air热重载工具
go install github.com/air-verse/air@latest

# 初始化Air配置
air init

# 启动热重载开发服务
air
```

现在访问 `http://localhost:9000` 即可看到应用运行效果。修改代码后会自动重新编译和重启！

## 🛠️ CLI脚手架工具

本项目还提供了一个强大的CLI脚手架工具，可以帮助您快速创建新的热重载项目！

### 安装CLI工具

```bash
go install github.com/yggai/aigo_hotreload@latest
```

### 使用CLI工具创建新项目
```bash
# 创建新项目
aigo_hotreload create my-new-api

# 进入项目目录
cd my-new-api

# 安装依赖并启动
go mod tidy
air
```

### CLI工具特性
- ✅ **一键创建**：自动生成完整项目结构
- ✅ **预配置热重载**：内置Air配置文件
- ✅ **标准化结构**：包含README、.gitignore等
- ✅ **即开即用**：生成后立即可运行



### 🌐 域名配置（生产环境）

#### 1. 域名解析
将您的域名（如：`testapi.zhangdapeng.com`）解析到服务器IP地址。

#### 2. Nginx反向代理配置

创建Nginx站点配置文件：
```bash
sudo nano /etc/nginx/sites-available/your-domain.com
```

配置内容：
```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:9000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        
        # WebSocket支持
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        
        # 超时设置
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
    }

    # 日志配置
    access_log /var/log/nginx/your-domain.com.access.log;
    error_log /var/log/nginx/your-domain.com.error.log;
}
```

#### 3. 启用站点
```bash
# 创建软链接
sudo ln -sf /etc/nginx/sites-available/your-domain.com /etc/nginx/sites-enabled/

# 测试配置
sudo nginx -t

# 重新加载Nginx
sudo systemctl reload nginx
```

### 🔒 HTTPS SSL证书配置

#### 1. 安装Certbot
```bash
sudo apt update
sudo apt install -y certbot python3-certbot-nginx
```

#### 2. 申请SSL证书
```bash
sudo certbot --nginx -d your-domain.com
```

#### 3. 验证HTTPS
访问 `https://your-domain.com` 确认SSL证书配置成功。

### 📝 核心代码文件

#### go.mod
```go
module github.com/yggai/aigo_hotreload

go 1.24

require github.com/gin-gonic/gin v1.10.1
```

#### main.go
```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建Gin路由器
	r := gin.Default()

	// 启动信息
	fmt.Println("🚀 正在启动Gin服务器...")
	fmt.Println("📡 服务器将在 http://localhost:9000 启动")

	// 根路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "热更新测试成功！代码已自动重载",
			"status":    "success",
			"timestamp": "2024-01-01",
		})
	})

	// Hello路由
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "Hello, World!",
			"greeting": "欢迎来到Gin世界！",
			"version":  "v1.0.0",
		})
	})

	// 启动服务器
	r.Run(":9000")
}
```

### 🔧 常用命令

#### 开发命令
```bash
# 启动热重载开发
air

# 停止热重载服务
Ctrl+C  # 或者
killall air

# 重新初始化Air配置
air init

# 直接运行（不使用热重载）
go run main.go
```

#### 生产环境命令
```bash
# 编译生产版本
go build -o app main.go

# 运行生产版本
./app

# 后台运行
nohup ./app &
```

### 🔍 故障排除

#### 常见问题

1. **端口被占用**
   ```bash
   # 查看端口占用
   lsof -i :9000
   
   # 杀死占用进程
   kill -9 <PID>
   ```

2. **Nginx配置错误**
   ```bash
   # 测试配置
   sudo nginx -t
   
   # 查看错误日志
   sudo tail -f /var/log/nginx/error.log
   ```

3. **SSL证书问题**
   ```bash
   # 检查证书状态
   sudo certbot certificates
   
   # 手动续期
   sudo certbot renew
   ```

### 📊 性能优化

- 生产环境建议使用 `gin.SetMode(gin.ReleaseMode)` 关闭调试模式
- 配置适当的Nginx缓存策略
- 使用反向代理负载均衡（多实例部署）

### 📄 许可证

本项目采用个人研究项目协议，允许学习和研究使用，商业使用需要作者书面授权。

### 👨‍💻 作者信息

- **作者**: 源滚滚AI编程
- **邮箱**: 1156956626@qq.com
- **电话**: 18010070052

---

**Happy Coding! 🎉**
