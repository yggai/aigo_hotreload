# Go 热重载开发模板 (aigo_hotreload)

> 🚀 源滚滚AI编程开发的Go语言热重载模板项目，专为云服务器开发环境设计，支持本地和域名热更新，大幅提升开发效率！

[中文](README.md) | [English](README_EN.md)

---

### 📋 项目简介

本项目是一个完整的Go语言热重载开发模板，集成了以下特性：
- ✅ 本地热重载开发
- ✅ 域名反向代理配置
- ✅ HTTPS SSL证书自动配置（自动检测并安装certbot）
- ✅ Nginx反向代理支持
- ✅ 生产环境部署指南
- ✅ 智能依赖管理（自动适配不同Linux发行版）

### 🛠️ 技术栈

- **后端框架**: Gin (Go)
- **热重载工具**: Air
- **反向代理**: Nginx
- **SSL证书**: Let's Encrypt (免费)
- **部署环境**: Linux云服务器

### 🚀 核心优势

#### 智能化部署
- **零配置启动**：一键创建项目，无需手动配置
- **智能依赖管理**：自动检测并安装所需工具（certbot、nginx等）
- **跨平台兼容**：支持Ubuntu、Debian、CentOS、Fedora等主流Linux发行版

#### 开发效率提升
- **热重载开发**：代码修改后自动重新编译和重启
- **标准化结构**：统一的项目结构和最佳实践
- **即开即用**：生成后立即可运行，无需额外配置

#### 生产环境就绪
- **SSL证书自动化**：一键申请和配置HTTPS证书
- **nginx配置模板**：预配置的反向代理和负载均衡
- **完整部署流程**：从开发到生产的完整解决方案

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

现在访问 `http://localhost:8888` 即可看到应用运行效果。修改代码后会自动重新编译和重启！

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

# 生成nginx配置（可选）
aigo_hotreload nginx your-domain.com ./my-new-api 8888
```

### CLI工具完整功能

#### 创建项目
```bash
# 创建新的热重载项目
aigo_hotreload create <project-name>
```

#### 生成nginx配置
```bash
# 为指定域名生成nginx配置文件
aigo_hotreload nginx <domain> <project-path> [port]

# 示例
aigo_hotreload nginx api.example.com ./my-api 8888
```

#### 查看帮助
```bash
# 显示帮助信息
aigo_hotreload help

# 显示版本信息
aigo_hotreload version
```

#### 生成的项目结构
```
my-api/
├── main.go              # 主程序文件
├── go.mod              # Go模块依赖
├── .air.toml           # Air热重载配置
├── .gitignore          # Git忽略文件
├── README.md           # 项目说明
├── config/             # nginx配置文件目录
│   ├── your-domain.com # nginx配置文件
│   └── setup-nginx.sh  # nginx配置脚本
└── scripts/            # 脚本目录
    └── apply-ssl.sh    # SSL证书申请脚本
```

### CLI工具特性
- ✅ **一键创建**：自动生成完整项目结构
- ✅ **预配置热重载**：内置Air配置文件
- ✅ **标准化结构**：包含README、.gitignore等
- ✅ **即开即用**：生成后立即可运行
- ✅ **nginx配置**：自动生成nginx配置文件和脚本
- ✅ **SSL证书**：自动生成SSL证书申请脚本（智能检测并安装certbot）
- ✅ **域名部署**：支持一键配置域名和HTTPS
- ✅ **智能依赖管理**：自动适配Ubuntu/Debian/CentOS/Fedora等系统



### 🌐 域名配置（生产环境）

#### 1. 使用CLI工具自动生成nginx配置

```bash
# 创建新项目
aigo_hotreload create my-api

# 生成nginx配置文件
aigo_hotreload nginx your-domain.com ./my-api 8888
```

#### 2. 手动配置nginx

**域名解析**
将您的域名（如：`testapi.zhangdapeng.com`）解析到服务器IP地址。

**创建Nginx站点配置文件**
```bash
sudo nano /etc/nginx/sites-available/your-domain.com
```

**配置内容**
```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:8888;
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

**启用站点**
```bash
# 创建软链接
sudo ln -sf /etc/nginx/sites-available/your-domain.com /etc/nginx/sites-enabled/

# 测试配置
sudo nginx -t

# 重新加载Nginx
sudo systemctl reload nginx
```

#### 3. 使用自动配置脚本

CLI工具生成的项目包含自动配置脚本：

```bash
# 进入项目目录
cd my-api

# 编辑nginx配置文件
vim config/your-domain.com

# 运行nginx配置脚本
chmod +x config/setup-nginx.sh
./config/setup-nginx.sh your-domain.com 8888
```

### 🔒 HTTPS SSL证书配置

#### 1. 使用CLI工具自动申请SSL证书

```bash
# 进入项目目录
cd my-api

# 申请SSL证书（脚本会自动检测并安装certbot）
chmod +x scripts/apply-ssl.sh
./scripts/apply-ssl.sh your-domain.com
```

**智能依赖管理示例**：
```bash
# 脚本会自动执行以下步骤：
# 1. 检查certbot是否已安装
# 2. 如果未安装，根据系统类型自动选择包管理器：
#    - Ubuntu/Debian: apt install -y certbot python3-certbot-nginx
#    - CentOS/RHEL: yum install -y certbot python3-certbot-nginx
#    - Fedora: dnf install -y certbot python3-certbot-nginx
# 3. 验证安装是否成功
# 4. 申请SSL证书
# 5. 配置nginx并重新加载
```

#### 2. 手动申请SSL证书

**安装Certbot**
```bash
sudo apt update
sudo apt install -y certbot python3-certbot-nginx
```

**申请SSL证书**
```bash
sudo certbot --nginx -d your-domain.com
```

**验证HTTPS**
访问 `https://your-domain.com` 确认SSL证书配置成功。

#### 3. 智能依赖管理特性

CLI工具生成的SSL申请脚本具有以下智能特性：
- 🔍 **自动检测certbot**：检查是否已安装certbot
- 🚀 **智能安装**：根据系统类型自动选择包管理器
  - Ubuntu/Debian：使用 `apt install`
  - CentOS/RHEL：使用 `yum install`
  - Fedora：使用 `dnf install`
- ✅ **安装验证**：验证安装是否成功
- 📋 **故障排除**：提供详细的错误诊断和解决建议

#### 4. SSL证书管理

```bash
# 查看证书状态
sudo certbot certificates

# 手动续期
sudo certbot renew

# 删除证书
sudo certbot delete --cert-name your-domain.com
```

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
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建Gin路由器
	r := gin.Default()

	// 启动信息
	fmt.Println("正在启动gin服务器...")
	fmt.Println("服务器将在 http://localhost:8888 启动")

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

	// 启动服务器
	r.Run(":8888")
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
   lsof -i :8888
   
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
