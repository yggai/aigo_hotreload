server {
    server_name testapi.zhangdapeng.com;

    location / {
        proxy_pass http://localhost:9000;
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
    access_log /var/log/nginx/testapi.zhangdapeng.com.access.log;
    error_log /var/log/nginx/testapi.zhangdapeng.com.error.log;

    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/testapi.zhangdapeng.com/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/testapi.zhangdapeng.com/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

}
server {
    if ($host = testapi.zhangdapeng.com) {
        return 301 https://$host$request_uri;
    } # managed by Certbot


    listen 80;
    server_name testapi.zhangdapeng.com;
    return 404; # managed by Certbot


}