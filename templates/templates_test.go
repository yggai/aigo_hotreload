package templates

import (
	"fmt"
	"strings"
	"testing"
)

// TestGoModTemplate 测试go.mod模板
func TestGoModTemplate(t *testing.T) {
	// 测试模板格式化
	projectName := "test-project"
	formatted := fmt.Sprintf(GoModTemplate, projectName)
	
	// 验证模板内容
	if !strings.Contains(formatted, "module "+projectName) {
		t.Errorf("go.mod模板应该包含模块名: %s", projectName)
	}
	
	if !strings.Contains(formatted, "go 1.24") {
		t.Errorf("go.mod模板应该包含Go版本")
	}
	
	if !strings.Contains(formatted, "github.com/gin-gonic/gin") {
		t.Errorf("go.mod模板应该包含gin依赖")
	}
	
	if !strings.Contains(formatted, "v1.10.1") {
		t.Errorf("go.mod模板应该包含gin版本")
	}
}

// TestMainGoTemplate 测试main.go模板
func TestMainGoTemplate(t *testing.T) {
	// 验证模板内容
	if !strings.Contains(MainGoTemplate, "package main") {
		t.Errorf("main.go模板应该包含package main")
	}
	
	if !strings.Contains(MainGoTemplate, "import") {
		t.Errorf("main.go模板应该包含import语句")
	}
	
	if !strings.Contains(MainGoTemplate, "gin.Default()") {
		t.Errorf("main.go模板应该包含gin.Default()")
	}
	
	if !strings.Contains(MainGoTemplate, ":8888") {
		t.Errorf("main.go模板应该包含端口8888")
	}
	
	if !strings.Contains(MainGoTemplate, "r.GET(\"/\"") {
		t.Errorf("main.go模板应该包含根路由")
	}
	
	if !strings.Contains(MainGoTemplate, "r.GET(\"/hello\"") {
		t.Errorf("main.go模板应该包含hello路由")
	}
	
	if !strings.Contains(MainGoTemplate, "r.GET(\"/health\"") {
		t.Errorf("main.go模板应该包含health路由")
	}
	
	if !strings.Contains(MainGoTemplate, "/api/v1") {
		t.Errorf("main.go模板应该包含API路由组")
	}
	
	if !strings.Contains(MainGoTemplate, "/users") {
		t.Errorf("main.go模板应该包含users路由")
	}
}

// TestAirTomlTemplate 测试.air.toml模板
func TestAirTomlTemplate(t *testing.T) {
	// 验证模板内容
	if !strings.Contains(AirTomlTemplate, "[build]") {
		t.Errorf(".air.toml模板应该包含[build]配置")
	}
	
	if !strings.Contains(AirTomlTemplate, "cmd = \"go build -o ./tmp/main .\"") {
		t.Errorf(".air.toml模板应该包含构建命令")
	}
	
	if !strings.Contains(AirTomlTemplate, "bin = \"./tmp/main\"") {
		t.Errorf(".air.toml模板应该包含二进制文件路径")
	}
	
	if !strings.Contains(AirTomlTemplate, "include_ext = [\"go\", \"tpl\", \"tmpl\", \"html\"]") {
		t.Errorf(".air.toml模板应该包含包含的文件扩展名")
	}
	
	if !strings.Contains(AirTomlTemplate, "exclude_dir = [\"assets\", \"tmp\", \"vendor\"]") {
		t.Errorf(".air.toml模板应该包含排除的目录")
	}
	
	if !strings.Contains(AirTomlTemplate, "delay = 1000") {
		t.Errorf(".air.toml模板应该包含延迟配置")
	}
}

// TestGitignoreTemplate 测试.gitignore模板
func TestGitignoreTemplate(t *testing.T) {
	// 验证模板内容
	if !strings.Contains(GitignoreTemplate, "*.exe") {
		t.Errorf(".gitignore模板应该包含*.exe")
	}
	
	if !strings.Contains(GitignoreTemplate, "*.exe~") {
		t.Errorf(".gitignore模板应该包含*.exe~")
	}
	
	if !strings.Contains(GitignoreTemplate, "*.dll") {
		t.Errorf(".gitignore模板应该包含*.dll")
	}
	
	if !strings.Contains(GitignoreTemplate, "*.so") {
		t.Errorf(".gitignore模板应该包含*.so")
	}
	
	if !strings.Contains(GitignoreTemplate, "*.dylib") {
		t.Errorf(".gitignore模板应该包含*.dylib")
	}
	
	if !strings.Contains(GitignoreTemplate, "*.test") {
		t.Errorf(".gitignore模板应该包含*.test")
	}
	
	if !strings.Contains(GitignoreTemplate, "*.out") {
		t.Errorf(".gitignore模板应该包含*.out")
	}
	
	if !strings.Contains(GitignoreTemplate, "tmp/") {
		t.Errorf(".gitignore模板应该包含tmp/目录")
	}
	
	if !strings.Contains(GitignoreTemplate, ".air.toml") {
		t.Errorf(".gitignore模板应该包含.air.toml")
	}
}

// TestReadmeTemplate 测试README.md模板
func TestReadmeTemplate(t *testing.T) {
	// 测试模板格式化
	projectName := "test-project"
	formatted := fmt.Sprintf(ReadmeTemplate, projectName, projectName)
	
	// 验证模板内容
	if !strings.Contains(formatted, "# "+projectName) {
		t.Errorf("README模板应该包含项目标题: %s", projectName)
	}
	
	if !strings.Contains(formatted, "go mod tidy") {
		t.Errorf("README模板应该包含go mod tidy命令")
	}
	
	if !strings.Contains(formatted, "air") {
		t.Errorf("README模板应该包含air命令")
	}
	
	if !strings.Contains(formatted, "http://localhost:8888") {
		t.Errorf("README模板应该包含访问地址")
	}
	
	if !strings.Contains(formatted, "域名配置") {
		t.Errorf("README模板应该包含域名配置说明")
	}
	
	if !strings.Contains(formatted, "HTTPS SSL证书") {
		t.Errorf("README模板应该包含SSL证书说明")
	}
	
	if !strings.Contains(formatted, "nginx配置") {
		t.Errorf("README模板应该包含nginx配置说明")
	}
}

// TestNginxHTTPTemplateInTemplates 测试nginx HTTP模板
func TestNginxHTTPTemplateInTemplates(t *testing.T) {
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

// TestNginxHTTPSTemplateInTemplates 测试nginx HTTPS模板
func TestNginxHTTPSTemplateInTemplates(t *testing.T) {
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
}

// TestNginxSetupScriptTemplateInTemplates 测试nginx设置脚本模板
func TestNginxSetupScriptTemplateInTemplates(t *testing.T) {
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
}

// TestCertbotScriptTemplateInTemplates 测试certbot脚本模板
func TestCertbotScriptTemplateInTemplates(t *testing.T) {
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
}

// TestTemplateVariables 测试模板变量替换
func TestTemplateVariables(t *testing.T) {
	// 测试go.mod模板变量替换
	projectName := "my-test-project"
	formatted := fmt.Sprintf(GoModTemplate, projectName)
	
	if !strings.Contains(formatted, "module "+projectName) {
		t.Errorf("go.mod模板应该正确替换项目名: %s", projectName)
	}
	
	// 测试nginx模板变量替换
	domain := "example.com"
	port := "8080"
	nginxConfig := fmt.Sprintf(NginxHTTPTemplate, domain, port, domain, domain)
	
	if !strings.Contains(nginxConfig, "server_name "+domain) {
		t.Errorf("nginx模板应该正确替换域名: %s", domain)
	}
	
	if !strings.Contains(nginxConfig, "proxy_pass http://localhost:"+port) {
		t.Errorf("nginx模板应该正确替换端口: %s", port)
	}
}

// TestTemplateSpecialCharacters 测试模板中的特殊字符
func TestTemplateSpecialCharacters(t *testing.T) {
	// 测试包含特殊字符的项目名
	projectName := "test-project-with-special-chars_123"
	formatted := fmt.Sprintf(GoModTemplate, projectName)
	
	if !strings.Contains(formatted, "module "+projectName) {
		t.Errorf("go.mod模板应该正确处理特殊字符项目名: %s", projectName)
	}
	
	// 测试包含特殊字符的域名
	domain := "test.example.com"
	port := "8888"
	nginxConfig := fmt.Sprintf(NginxHTTPTemplate, domain, port, domain, domain)
	
	if !strings.Contains(nginxConfig, "server_name "+domain) {
		t.Errorf("nginx模板应该正确处理特殊字符域名: %s", domain)
	}
}

// TestTemplateEmptyValues 测试模板空值处理
func TestTemplateEmptyValues(t *testing.T) {
	// 测试空项目名
	projectName := ""
	formatted := fmt.Sprintf(GoModTemplate, projectName)
	
	if !strings.Contains(formatted, "module ") {
		t.Error("go.mod模板应该能处理空项目名")
	}
	
	// 测试空域名
	domain := ""
	port := "8888"
	nginxConfig := fmt.Sprintf(NginxHTTPTemplate, domain, port, domain, domain)
	
	if !strings.Contains(nginxConfig, "server_name ") {
		t.Error("nginx模板应该能处理空域名")
	}
}

// TestTemplateLargeValues 测试模板大值处理
func TestTemplateLargeValues(t *testing.T) {
	// 测试长项目名
	projectName := strings.Repeat("a", 100)
	formatted := fmt.Sprintf(GoModTemplate, projectName)
	
	if !strings.Contains(formatted, "module "+projectName) {
		t.Errorf("go.mod模板应该能处理长项目名")
	}
	
	// 测试长域名
	domain := strings.Repeat("a", 50) + ".com"
	port := "8888"
	nginxConfig := fmt.Sprintf(NginxHTTPTemplate, domain, port, domain, domain)
	
	if !strings.Contains(nginxConfig, "server_name "+domain) {
		t.Errorf("nginx模板应该能处理长域名")
	}
}

// TestTemplateUnicode 测试模板Unicode字符
func TestTemplateUnicode(t *testing.T) {
	// 测试包含Unicode字符的项目名
	projectName := "测试项目-🚀-🎉"
	formatted := fmt.Sprintf(GoModTemplate, projectName)
	
	if !strings.Contains(formatted, "module "+projectName) {
		t.Errorf("go.mod模板应该能处理Unicode项目名: %s", projectName)
	}
	
	// 测试包含Unicode字符的域名
	domain := "测试域名.example.com"
	port := "8888"
	nginxConfig := fmt.Sprintf(NginxHTTPTemplate, domain, port, domain, domain)
	
	if !strings.Contains(nginxConfig, "server_name "+domain) {
		t.Errorf("nginx模板应该能处理Unicode域名: %s", domain)
	}
}

// TestTemplateFormatting 测试模板格式化
func TestTemplateFormatting(t *testing.T) {
	// 测试go.mod模板格式化
	projectName := "test-project"
	formatted := fmt.Sprintf(GoModTemplate, projectName)
	
	// 验证基本结构
	expectedLines := []string{
		"module " + projectName,
		"go 1.24",
		"github.com/gin-gonic/gin",
	}
	
	for _, line := range expectedLines {
		if !strings.Contains(formatted, line) {
			t.Errorf("go.mod模板应该包含: %s", line)
		}
	}
	
	// 测试main.go模板格式化
	if !strings.Contains(MainGoTemplate, "package main") {
		t.Error("main.go模板应该包含package main")
	}
	
	if !strings.Contains(MainGoTemplate, "import") {
		t.Error("main.go模板应该包含import语句")
	}
	
	if !strings.Contains(MainGoTemplate, "func main()") {
		t.Error("main.go模板应该包含main函数")
	}
}

// TestTemplateConsistency 测试模板一致性
func TestTemplateConsistency(t *testing.T) {
	// 测试端口号一致性
	if !strings.Contains(MainGoTemplate, ":8888") {
		t.Error("main.go模板应该使用端口8888")
	}
	
	if !strings.Contains(NginxHTTPTemplate, "8888") {
		t.Error("nginx模板应该使用端口8888")
	}
	
	// 测试路由一致性
	if !strings.Contains(MainGoTemplate, "/health") {
		t.Error("main.go模板应该包含health路由")
	}
	
	if !strings.Contains(MainGoTemplate, "/api/v1") {
		t.Error("main.go模板应该包含API路由")
	}
}

// TestTemplateSecurity 测试模板安全性
func TestTemplateSecurity(t *testing.T) {
	// 测试XSS防护
	maliciousProjectName := "<script>alert('xss')</script>"
	formatted := fmt.Sprintf(GoModTemplate, maliciousProjectName)
	
	// 验证恶意脚本没有被执行（只是作为字符串）
	if !strings.Contains(formatted, maliciousProjectName) {
		t.Error("模板应该正确处理恶意输入")
	}
	
	// 测试SQL注入防护
	sqlInjectionDomain := "'; DROP TABLE users; --"
	port := "8888"
	nginxConfig := fmt.Sprintf(NginxHTTPTemplate, sqlInjectionDomain, port, sqlInjectionDomain, sqlInjectionDomain)
	
	if !strings.Contains(nginxConfig, sqlInjectionDomain) {
		t.Error("模板应该正确处理SQL注入尝试")
	}
}

// TestTemplatePerformance 测试模板性能
func TestTemplatePerformance(t *testing.T) {
	// 测试大量格式化操作
	projectName := "performance-test-project"
	
	for i := 0; i < 1000; i++ {
		formatted := fmt.Sprintf(GoModTemplate, projectName)
		if !strings.Contains(formatted, "module "+projectName) {
			t.Errorf("模板格式化失败，迭代 %d", i)
		}
	}
}

// TestTemplateEdgeCases 测试模板边界情况
func TestTemplateEdgeCases(t *testing.T) {
	// 测试非常长的项目名
	veryLongProjectName := strings.Repeat("a", 1000)
	formatted := fmt.Sprintf(GoModTemplate, veryLongProjectName)
	
	if !strings.Contains(formatted, "module "+veryLongProjectName) {
		t.Error("模板应该能处理非常长的项目名")
	}
	
	// 测试包含换行符的项目名
	projectNameWithNewlines := "test\nproject\nname"
	formatted = fmt.Sprintf(GoModTemplate, projectNameWithNewlines)
	
	if !strings.Contains(formatted, "module "+projectNameWithNewlines) {
		t.Error("模板应该能处理包含换行符的项目名")
	}
	
	// 测试包含制表符的项目名
	projectNameWithTabs := "test\tproject\tname"
	formatted = fmt.Sprintf(GoModTemplate, projectNameWithTabs)
	
	if !strings.Contains(formatted, "module "+projectNameWithTabs) {
		t.Error("模板应该能处理包含制表符的项目名")
	}
}

// TestTemplateEncoding 测试模板编码
func TestTemplateEncoding(t *testing.T) {
	// 测试UTF-8编码
	chineseProjectName := "中文项目名"
	formatted := fmt.Sprintf(GoModTemplate, chineseProjectName)
	
	if !strings.Contains(formatted, "module "+chineseProjectName) {
		t.Errorf("模板应该正确处理UTF-8编码: %s", chineseProjectName)
	}
	
	// 测试emoji
	emojiProjectName := "🚀🎉💻"
	formatted = fmt.Sprintf(GoModTemplate, emojiProjectName)
	
	if !strings.Contains(formatted, "module "+emojiProjectName) {
		t.Errorf("模板应该正确处理emoji: %s", emojiProjectName)
	}
}

// TestTemplateValidation 测试模板验证
func TestTemplateValidation(t *testing.T) {
	// 验证所有模板都不为空
	templates := []struct {
		name     string
		template string
	}{
		{"GoModTemplate", GoModTemplate},
		{"MainGoTemplate", MainGoTemplate},
		{"AirTomlTemplate", AirTomlTemplate},
		{"GitignoreTemplate", GitignoreTemplate},
		{"ReadmeTemplate", ReadmeTemplate},
		{"NginxHTTPTemplate", NginxHTTPTemplate},
		{"NginxHTTPSTemplate", NginxHTTPSTemplate},
		{"NginxSetupScriptTemplate", NginxSetupScriptTemplate},
		{"CertbotScriptTemplate", CertbotScriptTemplate},
	}
	
	for _, tc := range templates {
		if tc.template == "" {
			t.Errorf("模板 %s 不能为空", tc.name)
		}
		
		// 检查模板长度
		if len(tc.template) < 10 {
			t.Errorf("模板 %s 内容太短: %d 字符", tc.name, len(tc.template))
		}
		
		// 检查模板是否包含必要的占位符
		if tc.name == "GoModTemplate" && !strings.Contains(tc.template, "%s") {
			t.Errorf("模板 %s 应该包含 %%s 占位符", tc.name)
		}
		
		if strings.Contains(tc.name, "Nginx") && !strings.Contains(tc.template, "%s") {
			t.Errorf("nginx模板 %s 应该包含 %%s 占位符", tc.name)
		}
	}
}

// TestTemplateIntegration 测试模板集成
func TestTemplateIntegration(t *testing.T) {
	// 测试所有模板的集成使用
	projectName := "integration-test-project"
	domain := "integration.example.com"
	port := "8888"
	
	// 测试go.mod模板
	goModContent := fmt.Sprintf(GoModTemplate, projectName)
	if !strings.Contains(goModContent, "module "+projectName) {
		t.Error("go.mod模板集成测试失败")
	}
	
	// 测试nginx模板
	nginxContent := fmt.Sprintf(NginxHTTPTemplate, domain, port, domain, domain)
	if !strings.Contains(nginxContent, "server_name "+domain) {
		t.Error("nginx模板集成测试失败")
	}
	
	// 测试README模板
	readmeContent := fmt.Sprintf(ReadmeTemplate, projectName, projectName)
	if !strings.Contains(readmeContent, projectName) {
		t.Error("README模板集成测试失败")
	}
}

// TestTemplateConcurrency 测试模板并发使用
func TestTemplateConcurrency(t *testing.T) {
	// 测试并发格式化模板
	done := make(chan bool, 10)
	
	for i := 0; i < 10; i++ {
		go func(id int) {
			projectName := fmt.Sprintf("concurrent-project-%d", id)
			formatted := fmt.Sprintf(GoModTemplate, projectName)
			
			if !strings.Contains(formatted, "module "+projectName) {
				t.Errorf("并发模板格式化失败: %s", projectName)
			}
			
			done <- true
		}(i)
	}
	
	// 等待所有goroutine完成
	for i := 0; i < 10; i++ {
		<-done
	}
} 