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
├── main.go              # Main program file
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

Now visit `http://localhost:9000` to see your application. Code changes will automatically recompile and restart!

## 🌐 Domain Configuration (Production)

### 1. DNS Setup
Point your domain (e.g., `testapi.zhangdapeng.com`) to your server IP address.

### 2. Nginx Reverse Proxy

Create Nginx site configuration:
```bash
sudo nano /etc/nginx/sites-available/your-domain.com
```

Configuration content:
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

### 3. Enable Site
```bash
# Create symbolic link
sudo ln -sf /etc/nginx/sites-available/your-domain.com /etc/nginx/sites-enabled/

# Test configuration
sudo nginx -t

# Reload Nginx
sudo systemctl reload nginx
```

## 🔒 HTTPS SSL Certificate

### 1. Install Certbot
```bash
sudo apt update
sudo apt install -y certbot python3-certbot-nginx
```

### 2. Obtain SSL Certificate
```bash
sudo certbot --nginx -d your-domain.com
```

### 3. Verify HTTPS
Visit `https://your-domain.com` to confirm SSL certificate is working.

## 📝 Core Code Files

### go.mod
```go
module github.com/yggai/aigo_hotreload

go 1.24.4

require github.com/gin-gonic/gin v1.10.1
```

### main.go
```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create Gin router
	r := gin.Default()

	// Startup information
	fmt.Println("🚀 Starting Gin server...")
	fmt.Println("📡 Server will start at http://localhost:9000")

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
		c.JSON(200, gin.H{
			"message":  "Hello, World!",
			"greeting": "Welcome to Gin world!",
			"version":  "v1.0.0",
		})
	})

	// Start server
	r.Run(":9000")
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
   lsof -i :9000
   
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