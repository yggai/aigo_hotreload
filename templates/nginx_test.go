package templates

import (
	"fmt"
	"strings"
	"testing"
)

// TestNginxHTTPTemplate 测试nginx HTTP模板
func TestNginxHTTPTemplate(t *testing.T) {
	// 测试模板格式化
	domain := "test.example.com"
	port := "8888"
	formatted := fmt.Sprintf(NginxHTTPTemplate, domain, port, domain, domain)
	
	// 验证模板内容
	if !strings.Contains(formatted, "server_name "+domain) {
		t.Errorf("nginx模板应该包含server_name: %s", domain)
	}
	
	if !strings.Contains(formatted, "proxy_pass http://localhost:"+port) {
		t.Errorf("nginx模板应该包含proxy_pass: %s", port)
	}
	
	if !strings.Contains(formatted, "listen 80") {
		t.Errorf("nginx模板应该包含listen 80")
	}
	
	if !strings.Contains(formatted, "proxy_set_header Host $host") {
		t.Errorf("nginx模板应该包含Host头设置")
	}
	
	if !strings.Contains(formatted, "proxy_set_header X-Real-IP $remote_addr") {
		t.Errorf("nginx模板应该包含X-Real-IP头设置")
	}
	
	if !strings.Contains(formatted, "proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for") {
		t.Errorf("nginx模板应该包含X-Forwarded-For头设置")
	}
	
	if !strings.Contains(formatted, "proxy_set_header X-Forwarded-Proto $scheme") {
		t.Errorf("nginx模板应该包含X-Forwarded-Proto头设置")
	}
	
	if !strings.Contains(formatted, "proxy_http_version 1.1") {
		t.Errorf("nginx模板应该包含HTTP版本设置")
	}
	
	if !strings.Contains(formatted, "proxy_set_header Upgrade $http_upgrade") {
		t.Errorf("nginx模板应该包含Upgrade头设置")
	}
	
	if !strings.Contains(formatted, "proxy_set_header Connection \"upgrade\"") {
		t.Errorf("nginx模板应该包含Connection头设置")
	}
	
	if !strings.Contains(formatted, "proxy_connect_timeout 60s") {
		t.Errorf("nginx模板应该包含连接超时设置")
	}
	
	if !strings.Contains(formatted, "proxy_send_timeout 60s") {
		t.Errorf("nginx模板应该包含发送超时设置")
	}
	
	if !strings.Contains(formatted, "proxy_read_timeout 60s") {
		t.Errorf("nginx模板应该包含读取超时设置")
	}
	
	if !strings.Contains(formatted, "access_log /var/log/nginx/"+domain+".access.log") {
		t.Errorf("nginx模板应该包含访问日志设置")
	}
	
	if !strings.Contains(formatted, "error_log /var/log/nginx/"+domain+".error.log") {
		t.Errorf("nginx模板应该包含错误日志设置")
	}
}

// TestNginxHTTPSTemplate 测试nginx HTTPS模板
func TestNginxHTTPSTemplate(t *testing.T) {
	// 测试模板格式化
	domain := "test.example.com"
	port := "8888"
	formatted := fmt.Sprintf(NginxHTTPSTemplate, domain, port, domain, domain, domain, domain, domain, domain)
	
	// 验证模板内容
	if !strings.Contains(formatted, "server_name "+domain) {
		t.Errorf("nginx HTTPS模板应该包含server_name: %s", domain)
	}
	
	if !strings.Contains(formatted, "listen 443 ssl") {
		t.Errorf("nginx HTTPS模板应该包含listen 443 ssl")
	}
	
	if !strings.Contains(formatted, "ssl_certificate /etc/letsencrypt/live/"+domain+"/fullchain.pem") {
		t.Errorf("nginx HTTPS模板应该包含SSL证书路径")
	}
	
	if !strings.Contains(formatted, "ssl_certificate_key /etc/letsencrypt/live/"+domain+"/privkey.pem") {
		t.Errorf("nginx HTTPS模板应该包含SSL私钥路径")
	}
	
	if !strings.Contains(formatted, "proxy_pass http://localhost:"+port) {
		t.Errorf("nginx HTTPS模板应该包含proxy_pass: %s", port)
	}
	
	if !strings.Contains(formatted, "ssl_protocols TLSv1.2 TLSv1.3") {
		t.Errorf("nginx HTTPS模板应该包含SSL协议设置")
	}
	
	if !strings.Contains(formatted, "ssl_ciphers") {
		t.Errorf("nginx HTTPS模板应该包含SSL加密套件设置")
	}
	
	if !strings.Contains(formatted, "ssl_prefer_server_ciphers on") {
		t.Errorf("nginx HTTPS模板应该包含SSL服务器加密套件偏好设置")
	}
}

// TestNginxSetupScriptTemplate 测试nginx设置脚本模板
func TestNginxSetupScriptTemplate(t *testing.T) {
	// 验证模板内容
	if !strings.Contains(NginxSetupScriptTemplate, "#!/bin/bash") {
		t.Errorf("nginx设置脚本模板应该包含shebang")
	}
	
	if !strings.Contains(NginxSetupScriptTemplate, "nginx -t") {
		t.Errorf("nginx设置脚本模板应该包含nginx配置测试")
	}
	
	if !strings.Contains(NginxSetupScriptTemplate, "systemctl reload nginx") {
		t.Errorf("nginx设置脚本模板应该包含nginx重载命令")
	}
	
	if !strings.Contains(NginxSetupScriptTemplate, "ln -sf") {
		t.Errorf("nginx设置脚本模板应该包含软链接创建")
	}
	
	if !strings.Contains(NginxSetupScriptTemplate, "/etc/nginx/sites-available") {
		t.Errorf("nginx设置脚本模板应该包含sites-available路径")
	}
	
	if !strings.Contains(NginxSetupScriptTemplate, "/etc/nginx/sites-enabled") {
		t.Errorf("nginx设置脚本模板应该包含sites-enabled路径")
	}
	
	if !strings.Contains(NginxSetupScriptTemplate, "echo \"nginx配置完成\"") {
		t.Errorf("nginx设置脚本模板应该包含完成消息")
	}
	
	if !strings.Contains(NginxSetupScriptTemplate, "echo \"请访问 http://$DOMAIN\"") {
		t.Errorf("nginx设置脚本模板应该包含访问提示")
	}
}

// TestCertbotScriptTemplate 测试certbot脚本模板
func TestCertbotScriptTemplate(t *testing.T) {
	// 验证模板内容
	if !strings.Contains(CertbotScriptTemplate, "#!/bin/bash") {
		t.Errorf("certbot脚本模板应该包含shebang")
	}
	
	if !strings.Contains(CertbotScriptTemplate, "certbot") {
		t.Errorf("certbot脚本模板应该包含certbot命令")
	}
	
	if !strings.Contains(CertbotScriptTemplate, "nginx -t") {
		t.Errorf("certbot脚本模板应该包含nginx配置测试")
	}
	
	if !strings.Contains(CertbotScriptTemplate, "systemctl reload nginx") {
		t.Errorf("certbot脚本模板应该包含nginx重载命令")
	}
	
	if !strings.Contains(CertbotScriptTemplate, "apt install") {
		t.Errorf("certbot脚本模板应该包含apt安装命令")
	}
	
	if !strings.Contains(CertbotScriptTemplate, "yum install") {
		t.Errorf("certbot脚本模板应该包含yum安装命令")
	}
	
	if !strings.Contains(CertbotScriptTemplate, "dnf install") {
		t.Errorf("certbot脚本模板应该包含dnf安装命令")
	}
	
	if !strings.Contains(CertbotScriptTemplate, "Debian/Ubuntu") {
		t.Errorf("certbot脚本模板应该包含Debian/Ubuntu系统检测")
	}
	
	if !strings.Contains(CertbotScriptTemplate, "CentOS/RHEL") {
		t.Errorf("certbot脚本模板应该包含CentOS/RHEL系统检测")
	}
	
	if !strings.Contains(CertbotScriptTemplate, "Fedora") {
		t.Errorf("certbot脚本模板应该包含Fedora系统检测")
	}
	
	if !strings.Contains(CertbotScriptTemplate, "command -v certbot") {
		t.Errorf("certbot脚本模板应该包含certbot检查命令")
	}
	
	if !strings.Contains(CertbotScriptTemplate, "certbot --nginx -d") {
		t.Errorf("certbot脚本模板应该包含SSL证书申请命令")
	}
	
	if !strings.Contains(CertbotScriptTemplate, "python3-certbot-nginx") {
		t.Errorf("certbot脚本模板应该包含nginx插件")
	}
	
	if !strings.Contains(CertbotScriptTemplate, "echo \"正在为域名") {
		t.Errorf("certbot脚本模板应该包含进度提示")
	}
	
	if !strings.Contains(CertbotScriptTemplate, "echo \"✅ SSL证书申请成功\"") {
		t.Errorf("certbot脚本模板应该包含成功消息")
	}
	
	if !strings.Contains(CertbotScriptTemplate, "echo \"❌ SSL证书申请失败\"") {
		t.Errorf("certbot脚本模板应该包含失败消息")
	}
}

// TestNginxTemplateFormatting 测试nginx模板格式化
func TestNginxTemplateFormatting(t *testing.T) {
	// 测试各种格式化场景
	testCases := []struct {
		name     string
		template string
		args     []interface{}
		expected []string
	}{
		{
			name:     "nginx HTTP模板",
			template: NginxHTTPTemplate,
			args:     []interface{}{"test.example.com", "8888", "test.example.com", "test.example.com"},
			expected: []string{"server_name test.example.com", "proxy_pass http://localhost:8888", "listen 80"},
		},
		{
			name:     "nginx HTTPS模板",
			template: NginxHTTPSTemplate,
			args:     []interface{}{"test.example.com", "8888", "test.example.com", "test.example.com"},
			expected: []string{"server_name test.example.com", "listen 443 ssl", "ssl_certificate"},
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			formatted := fmt.Sprintf(tc.template, tc.args...)
			
			for _, expected := range tc.expected {
				if !strings.Contains(formatted, expected) {
					t.Errorf("模板应该包含: %s", expected)
				}
			}
		})
	}
}

// TestNginxTemplateSpecialCharacters 测试特殊字符处理
func TestNginxTemplateSpecialCharacters(t *testing.T) {
	// 测试包含特殊字符的域名
	domain := "test-domain_with.dots-and-dashes.com"
	port := "8888"
	formatted := fmt.Sprintf(NginxHTTPTemplate, domain, port, domain, domain)
	
	if !strings.Contains(formatted, "server_name "+domain) {
		t.Errorf("nginx模板应该正确处理特殊字符域名: %s", domain)
	}
	
	// 测试包含特殊字符的端口
	domain = "test.example.com"
	port = "8080"
	formatted = fmt.Sprintf(NginxHTTPTemplate, domain, port, domain, domain)
	
	if !strings.Contains(formatted, "proxy_pass http://localhost:"+port) {
		t.Errorf("nginx模板应该正确处理端口: %s", port)
	}
}

// TestNginxTemplateEmptyValues 测试空值处理
func TestNginxTemplateEmptyValues(t *testing.T) {
	// 测试空域名
	domain := ""
	port := "8888"
	formatted := fmt.Sprintf(NginxHTTPTemplate, domain, port, domain, domain)
	
	if !strings.Contains(formatted, "server_name ") {
		t.Errorf("nginx模板应该处理空域名")
	}
	
	// 测试空端口
	domain = "test.example.com"
	port = ""
	formatted = fmt.Sprintf(NginxHTTPTemplate, domain, port, domain, domain)
	
	if !strings.Contains(formatted, "proxy_pass http://localhost:") {
		t.Errorf("nginx模板应该处理空端口")
	}
}

// TestNginxTemplateConsistency 测试模板一致性
func TestNginxTemplateConsistency(t *testing.T) {
	// 验证HTTP和HTTPS模板的一致性
	domain := "test.example.com"
	port := "8888"
	
	httpFormatted := fmt.Sprintf(NginxHTTPTemplate, domain, port, domain, domain)
	httpsFormatted := fmt.Sprintf(NginxHTTPSTemplate, domain, port, domain, domain, domain, domain, domain, domain)
	
	// 两个模板都应该包含相同的server_name
	if !strings.Contains(httpFormatted, "server_name "+domain) {
		t.Errorf("HTTP模板应该包含server_name: %s", domain)
	}
	
	if !strings.Contains(httpsFormatted, "server_name "+domain) {
		t.Errorf("HTTPS模板应该包含server_name: %s", domain)
	}
	
	// 两个模板都应该包含相同的proxy_pass
	if !strings.Contains(httpFormatted, "proxy_pass http://localhost:"+port) {
		t.Errorf("HTTP模板应该包含proxy_pass: %s", port)
	}
	
	if !strings.Contains(httpsFormatted, "proxy_pass http://localhost:"+port) {
		t.Errorf("HTTPS模板应该包含proxy_pass: %s", port)
	}
}

// TestNginxTemplateCompleteness 测试模板完整性
func TestNginxTemplateCompleteness(t *testing.T) {
	// 验证所有必需的nginx模板都存在
	templates := []struct {
		name     string
		template string
	}{
		{"NginxHTTPTemplate", NginxHTTPTemplate},
		{"NginxHTTPSTemplate", NginxHTTPSTemplate},
		{"NginxSetupScriptTemplate", NginxSetupScriptTemplate},
		{"CertbotScriptTemplate", CertbotScriptTemplate},
	}
	
	for _, tc := range templates {
		if tc.template == "" {
			t.Errorf("模板 %s 不能为空", tc.name)
		}
		
		if len(tc.template) < 50 {
			t.Errorf("模板 %s 内容太短", tc.name)
		}
	}
}

// TestNginxTemplateSecurity 测试nginx模板安全性
func TestNginxTemplateSecurity(t *testing.T) {
	domain := "test.example.com"
	port := "8888"
	
	// 测试HTTP模板的安全设置
	httpFormatted := fmt.Sprintf(NginxHTTPTemplate, domain, port, domain, domain)
	
	// 验证包含安全相关的头设置
	securityHeaders := []string{
		"proxy_set_header Host $host",
		"proxy_set_header X-Real-IP $remote_addr",
		"proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for",
		"proxy_set_header X-Forwarded-Proto $scheme",
	}
	
	for _, header := range securityHeaders {
		if !strings.Contains(httpFormatted, header) {
			t.Errorf("HTTP模板应该包含安全头: %s", header)
		}
	}
	
	// 测试HTTPS模板的安全设置
	httpsFormatted := fmt.Sprintf(NginxHTTPSTemplate, domain, port, domain, domain, domain, domain, domain, domain)
	
	// 验证包含SSL安全设置
	sslSecuritySettings := []string{
		"ssl_protocols TLSv1.2 TLSv1.3",
		"ssl_prefer_server_ciphers on",
		"ssl_ciphers",
	}
	
	for _, setting := range sslSecuritySettings {
		if !strings.Contains(httpsFormatted, setting) {
			t.Errorf("HTTPS模板应该包含SSL安全设置: %s", setting)
		}
	}
}

// TestNginxTemplatePerformance 测试nginx模板性能设置
func TestNginxTemplatePerformance(t *testing.T) {
	domain := "test.example.com"
	port := "8888"
	
	// 测试HTTP模板的性能设置
	httpFormatted := fmt.Sprintf(NginxHTTPTemplate, domain, port, domain, domain)
	
	// 验证包含性能相关的设置
	performanceSettings := []string{
		"proxy_connect_timeout 60s",
		"proxy_send_timeout 60s",
		"proxy_read_timeout 60s",
		"proxy_http_version 1.1",
	}
	
	for _, setting := range performanceSettings {
		if !strings.Contains(httpFormatted, setting) {
			t.Errorf("HTTP模板应该包含性能设置: %s", setting)
		}
	}
}

// TestNginxTemplateWebSocket 测试nginx模板WebSocket支持
func TestNginxTemplateWebSocket(t *testing.T) {
	domain := "test.example.com"
	port := "8888"
	
	// 测试HTTP模板的WebSocket支持
	httpFormatted := fmt.Sprintf(NginxHTTPTemplate, domain, port, domain, domain)
	
	// 验证包含WebSocket相关的设置
	websocketSettings := []string{
		"proxy_http_version 1.1",
		"proxy_set_header Upgrade $http_upgrade",
		"proxy_set_header Connection \"upgrade\"",
	}
	
	for _, setting := range websocketSettings {
		if !strings.Contains(httpFormatted, setting) {
			t.Errorf("HTTP模板应该包含WebSocket设置: %s", setting)
		}
	}
}

// TestNginxTemplateLogging 测试nginx模板日志设置
func TestNginxTemplateLogging(t *testing.T) {
	domain := "test.example.com"
	port := "8888"
	
	// 测试HTTP模板的日志设置
	httpFormatted := fmt.Sprintf(NginxHTTPTemplate, domain, port, domain, domain)
	
	// 验证包含日志相关的设置
	loggingSettings := []string{
		"access_log /var/log/nginx/" + domain + ".access.log",
		"error_log /var/log/nginx/" + domain + ".error.log",
	}
	
	for _, setting := range loggingSettings {
		if !strings.Contains(httpFormatted, setting) {
			t.Errorf("HTTP模板应该包含日志设置: %s", setting)
		}
	}
} 