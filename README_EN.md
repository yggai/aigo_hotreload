# Go Hot-Reload Development Template (aigo_hotreload)

> 🚀 A complete Go hot-reload development template developed by YuanGunGun AI Programming, designed for cloud server development environments, supporting both local and domain hot-reload, greatly improving development efficiency!

[中文](README.md) | [English](README_EN.md)

---

## 📋 Project Overview

This is a complete Go hot-reload development template with the following features:
- ✅ Local hot-reload development
- ✅ Domain reverse proxy configuration
- ✅ Automatic HTTPS SSL certificate setup
- ✅ Nginx reverse proxy support
- ✅ Production deployment guide

## 🛠️ Tech Stack

- **Backend Framework**: Gin (Go)
- **Hot Reload Tool**: Air
- **Reverse Proxy**: Nginx
- **SSL Certificate**: Let's Encrypt (Free)
- **Deployment**: Linux Cloud Server

## 📁 Project Structure

```
aigo_hotreload/
├── main.go              # CLI tool main program
├── server.go            # Example server file
├── go.mod              # Go module dependencies
├── go.sum              # Dependency checksum file
├── .air.toml           # Air hot-reload configuration
├── config/             # Configuration directory
│   └── testapi.zhangdapeng.com  # Nginx site configuration
├── docs/               # Documentation directory
├── README.md           # Project documentation (Chinese)
└── README_EN.md        # Project documentation (English)
```

## 🚀 Quick Start

### Environment Requirements
- Go 1.24 or higher

### 1. Clone Repository
```bash
git clone https://github.com/yggai/aigo_hotreload.git
cd aigo_hotreload
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Local Development (Hot Reload)
```bash
# Install Air hot-reload tool
go install github.com/air-verse/air@latest

# Initialize Air configuration
air init

# Start hot-reload development server
air
```

Now visit `http://localhost:8888` to see your application. Code changes will automatically recompile and restart!

## 🛠️ CLI Scaffolding Tool

This project also provides a powerful CLI scaffolding tool to help you quickly create new hot-reload projects!

### Install CLI Tool

```bash
go install github.com/yggai/aigo_hotreload@latest
```

### Create New Project with CLI Tool
```bash
# Create new project
aigo_hotreload create my-new-api

# Enter project directory
cd my-new-api

# Install dependencies and start
go mod tidy
air

# Generate nginx configuration (optional)
aigo_hotreload nginx your-domain.com ./my-new-api 8888
```

### CLI Tool Complete Features

#### Create Project
```bash
# Create new hot-reload project
aigo_hotreload create <project-name>
```

#### Generate Nginx Configuration
```bash
# Generate nginx configuration for specified domain
aigo_hotreload nginx <domain> <project-path> [port]

# Example
aigo_hotreload nginx api.example.com ./my-api 8888
```

#### View Help
```bash
# Show help information
aigo_hotreload help

# Show version information
aigo_hotreload version
```

#### Generated Project Structure
```
my-api/
├── main.go              # Main program file
├── go.mod              # Go module dependencies
├── .air.toml           # Air hot-reload configuration
├── .gitignore          # Git ignore file
├── README.md           # Project documentation
├── config/             # Nginx configuration directory
│   ├── your-domain.com # Nginx configuration file
│   └── setup-nginx.sh  # Nginx setup script
└── scripts/            # Scripts directory
    └── apply-ssl.sh    # SSL certificate application script
```

### CLI Tool Features
- ✅ **One-click Creation**: Automatically generate complete project structure
- ✅ **Pre-configured Hot Reload**: Built-in Air configuration
- ✅ **Standardized Structure**: Includes README, .gitignore, etc.
- ✅ **Ready to Use**: Can run immediately after generation
- ✅ **Nginx Configuration**: Automatically generate nginx configuration files and scripts
- ✅ **SSL Certificate**: Automatically generate SSL certificate application scripts
- ✅ **Domain Deployment**: Support one-click domain and HTTPS configuration



## 🌐 Domain Configuration (Production)

### 1. Use CLI Tool to Auto-generate Nginx Configuration

```bash
# Create new project
aigo_hotreload create my-api

# Generate nginx configuration
aigo_hotreload nginx your-domain.com ./my-api 8888
```

### 2. Manual Nginx Configuration

**DNS Setup**
Point your domain (e.g., `testapi.zhangdapeng.com`) to your server IP address.

**Create Nginx Site Configuration**
```bash
sudo nano /etc/nginx/sites-available/your-domain.com
```

**Configuration Content**
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
        
        # WebSocket support
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        
        # Timeout settings
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
    }

    # Logging
    access_log /var/log/nginx/your-domain.com.access.log;
    error_log /var/log/nginx/your-domain.com.error.log;
}
```

**Enable Site**
```bash
# Create symbolic link
sudo ln -sf /etc/nginx/sites-available/your-domain.com /etc/nginx/sites-enabled/

# Test configuration
sudo nginx -t

# Reload Nginx
sudo systemctl reload nginx
```

### 3. Use Auto-configuration Scripts

Projects generated by CLI tool include auto-configuration scripts:

```bash
# Enter project directory
cd my-api

# Edit nginx configuration file
vim config/your-domain.com

# Run nginx setup script
chmod +x config/setup-nginx.sh
./config/setup-nginx.sh your-domain.com 8888
```

## 🔒 HTTPS SSL Certificate

### 1. Use CLI Tool to Auto-apply SSL Certificate

```bash
# Enter project directory
cd my-api

# Apply SSL certificate
chmod +x scripts/apply-ssl.sh
./scripts/apply-ssl.sh your-domain.com
```

### 2. Manual SSL Certificate Application

**Install Certbot**
```bash
sudo apt update
sudo apt install -y certbot python3-certbot-nginx
```

**Obtain SSL Certificate**
```bash
sudo certbot --nginx -d your-domain.com
```

**Verify HTTPS**
Visit `https://your-domain.com` to confirm SSL certificate is working.

### 3. SSL Certificate Management

```bash
# Check certificate status
sudo certbot certificates

# Manual renewal
sudo certbot renew

# Delete certificate
sudo certbot delete --cert-name your-domain.com
```

## 📝 Core Code Files

### go.mod
```go
module github.com/yggai/aigo_hotreload

go 1.24

require github.com/gin-gonic/gin v1.10.1
```

### main.go
```go
package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create Gin router
	r := gin.Default()

	// Startup information
	fmt.Println("Starting gin server...")
	fmt.Println("Server will start at http://localhost:8888")

	// Root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "Hot reload test successful! Code automatically reloaded",
			"status":    "success",
			"timestamp": "2024-01-01",
		})
	})

	// Hello route
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "Hello, World!",
			"greeting": "Welcome to gin world!",
		})
	})

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"time":   "2024-01-01 00:00:00",
		})
	})

	// API route group
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

	// Start server
	r.Run(":8888")
}
```

## 🔧 Common Commands

### Development Commands
```bash
# Start hot-reload development
air

# Stop hot-reload service
Ctrl+C  # or
killall air

# Reinitialize Air configuration
air init

# Run directly (without hot reload)
go run main.go
```

### Production Commands
```bash
# Build for production
go build -o app main.go

# Run production build
./app

# Run in background
nohup ./app &
```

## 🔍 Troubleshooting

### Common Issues

1. **Port Already in Use**
   ```bash
   # Check port usage
   lsof -i :8888
   
   # Kill process
   kill -9 <PID>
   ```

2. **Nginx Configuration Error**
   ```bash
   # Test configuration
   sudo nginx -t
   
   # Check error logs
   sudo tail -f /var/log/nginx/error.log
   ```

3. **SSL Certificate Issues**
   ```bash
   # Check certificate status
   sudo certbot certificates
   
   # Manual renewal
   sudo certbot renew
   ```

## 📊 Performance Optimization

- Use `gin.SetMode(gin.ReleaseMode)` to disable debug mode in production
- Configure appropriate Nginx caching strategies
- Use reverse proxy load balancing (multi-instance deployment)

## 📄 License

This project uses a personal research project agreement. Learning and research use is allowed, commercial use requires written authorization from the author.

## 👨‍💻 Author

- **Author**: 源滚滚AI编程 (YuanGunGun AI Programming)
- **Email**: 1156956626@qq.com
- **Phone**: 18010070052

---

**Happy Coding! 🎉**