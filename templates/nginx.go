package templates

// NginxHTTPTemplate HTTP版本的nginx配置文件模板
const NginxHTTPTemplate = `server {
    listen 80;
    server_name %s;

    location / {
        proxy_pass http://localhost:%s;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        
        # 支持WebSocket连接（如果需要）
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        
        # 超时设置
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
    }

    # 日志配置
    access_log /var/log/nginx/%s.access.log;
    error_log /var/log/nginx/%s.error.log;
}
`

// NginxHTTPSTemplate HTTPS版本的nginx配置文件模板（包含SSL证书）
const NginxHTTPSTemplate = `server {
    server_name %s;

    location / {
        proxy_pass http://localhost:%s;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        
        # 支持WebSocket连接（如果需要）
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        
        # 超时设置
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
    }

    # 日志配置
    access_log /var/log/nginx/%s.access.log;
    error_log /var/log/nginx/%s.error.log;

    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/%s/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/%s/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot
}

server {
    if ($host = %s) {
        return 301 https://$host$request_uri;
    } # managed by Certbot

    listen 80;
    server_name %s;
    return 404; # managed by Certbot
}
`

// CertbotScriptTemplate certbot申请SSL证书的脚本模板
const CertbotScriptTemplate = `#!/bin/bash
# SSL证书申请脚本
# 使用方法: ./apply-ssl.sh your-domain.com

if [ $# -eq 0 ]; then
    echo "请提供域名参数"
    echo "使用方法: ./apply-ssl.sh your-domain.com"
    exit 1
fi

DOMAIN=$1
NGINX_CONFIG="/etc/nginx/sites-enabled/$DOMAIN"

echo "正在为域名 $DOMAIN 申请SSL证书..."

# 检查nginx配置是否存在
if [ ! -f "$NGINX_CONFIG" ]; then
    echo "错误: nginx配置文件 $NGINX_CONFIG 不存在"
    echo "请先创建nginx配置文件并启用"
    exit 1
fi

# 检查certbot是否已安装
echo "检查certbot是否已安装..."
if ! command -v certbot &> /dev/null; then
    echo "certbot未安装，正在自动安装..."
    
    # 检测操作系统类型
    if [ -f /etc/debian_version ]; then
        # Debian/Ubuntu系统
        echo "检测到Debian/Ubuntu系统，使用apt安装..."
        apt update
        apt install -y certbot python3-certbot-nginx
    elif [ -f /etc/redhat-release ]; then
        # CentOS/RHEL系统
        echo "检测到CentOS/RHEL系统，使用yum安装..."
        yum install -y certbot python3-certbot-nginx
    elif command -v dnf &> /dev/null; then
        # Fedora系统
        echo "检测到Fedora系统，使用dnf安装..."
        dnf install -y certbot python3-certbot-nginx
    else
        echo "❌ 无法检测操作系统类型，请手动安装certbot："
        echo "Debian/Ubuntu: sudo apt install -y certbot python3-certbot-nginx"
        echo "CentOS/RHEL: sudo yum install -y certbot python3-certbot-nginx"
        echo "Fedora: sudo dnf install -y certbot python3-certbot-nginx"
        exit 1
    fi
    
    # 验证安装是否成功
    if ! command -v certbot &> /dev/null; then
        echo "❌ certbot安装失败，请手动安装后重试"
        exit 1
    else
        echo "✅ certbot安装成功"
    fi
else
    echo "✅ certbot已安装"
fi

# 申请SSL证书
echo "正在申请SSL证书..."
certbot --nginx -d $DOMAIN

# 检查证书申请是否成功
if [ $? -eq 0 ]; then
    echo "✅ SSL证书申请成功！"
    echo "现在您可以通过以下方式访问您的服务："
    echo "- HTTP:  http://$DOMAIN"
    echo "- HTTPS: https://$DOMAIN"
    
    # 测试nginx配置
    echo "正在测试nginx配置..."
    nginx -t && systemctl reload nginx
    
    if [ $? -eq 0 ]; then
        echo "✅ nginx配置更新成功！"
    else
        echo "❌ nginx配置更新失败，请检查配置"
    fi
else
    echo "❌ SSL证书申请失败"
    echo "请检查域名是否正确指向此服务器"
    echo "常见问题："
    echo "1. 域名DNS解析是否正确指向此服务器"
    echo "2. 80和443端口是否开放"
    echo "3. nginx是否正常运行"
fi
`

// NginxSetupScriptTemplate nginx配置脚本模板
const NginxSetupScriptTemplate = `#!/bin/bash
# nginx配置脚本
# 使用方法: ./setup-nginx.sh your-domain.com [port]

if [ $# -eq 0 ]; then
    echo "请提供域名参数"
    echo "使用方法: ./setup-nginx.sh your-domain.com [port]"
    exit 1
fi

DOMAIN=$1
PORT=${2:-8888}
CONFIG_FILE="config/$DOMAIN"
SITES_ENABLED="/etc/nginx/sites-enabled/$DOMAIN"

echo "正在为域名 $DOMAIN 配置nginx..."

# 检查配置文件是否存在
if [ ! -f "$CONFIG_FILE" ]; then
    echo "错误: 配置文件 $CONFIG_FILE 不存在"
    echo "请先运行项目生成器创建配置文件"
    exit 1
fi

# 创建软链接
echo "正在创建nginx软链接..."
ln -sf $(pwd)/$CONFIG_FILE $SITES_ENABLED

# 测试nginx配置
echo "正在测试nginx配置..."
nginx -t

if [ $? -eq 0 ]; then
    echo "✅ nginx配置测试通过"
    
    # 重新加载nginx配置
    echo "正在重新加载nginx配置..."
    systemctl reload nginx
    
    if [ $? -eq 0 ]; then
        echo "✅ nginx配置更新成功！"
        echo "现在您可以通过以下方式访问您的服务："
        echo "- HTTP: http://$DOMAIN"
        echo "- 本地: http://localhost:$PORT"
    else
        echo "❌ nginx配置更新失败"
    fi
else
    echo "❌ nginx配置测试失败，请检查配置文件"
fi
` 