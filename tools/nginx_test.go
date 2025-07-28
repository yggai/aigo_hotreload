package tools

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestNginxManagerCreation 测试NginxManager创建
func TestNginxManagerCreation(t *testing.T) {
	manager := NewNginxManager()
	if manager == nil {
		t.Error("NewNginxManager() 应该返回非空的NginxManager实例")
	}
}

// TestGenerateConfig 测试生成nginx配置
func TestGenerateConfig(t *testing.T) {
	manager := NewNginxManager()
	
	// 创建临时目录用于测试
	tempDir := t.TempDir()
	
	// 测试生成配置
	domain := "test.example.com"
	port := "8888"
	
	err := manager.GenerateConfig(domain, tempDir, port)
	if err != nil {
		t.Errorf("生成nginx配置失败: %v", err)
	}
	
	// 验证配置文件是否生成
	configFile := filepath.Join(tempDir, domain)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		t.Errorf("nginx配置文件应该被生成: %s", configFile)
	}
	
	// 验证配置文件内容
	content, err := os.ReadFile(configFile)
	if err != nil {
		t.Errorf("读取配置文件失败: %v", err)
	}
	
	contentStr := string(content)
	
	// 验证配置文件包含必要的配置项
	expectedItems := []string{
		"server_name " + domain,
		"proxy_pass http://localhost:" + port,
		"listen 80",
	}
	
	for _, item := range expectedItems {
		if !strings.Contains(contentStr, item) {
			t.Errorf("配置文件应该包含: %s", item)
		}
	}
}

// TestGenerateConfigWithEmptyPort 测试使用空端口生成配置
func TestGenerateConfigWithEmptyPort(t *testing.T) {
	manager := NewNginxManager()
	
	// 创建临时目录用于测试
	tempDir := t.TempDir()
	
	// 测试生成配置（空端口）
	domain := "test.example.com"
	port := ""
	
	err := manager.GenerateConfig(domain, tempDir, port)
	if err != nil {
		t.Errorf("生成nginx配置失败: %v", err)
	}
	
	// 验证配置文件是否生成
	configFile := filepath.Join(tempDir, domain)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		t.Errorf("nginx配置文件应该被生成: %s", configFile)
	}
}

// TestGenerateSetupScript 测试生成nginx设置脚本
func TestGenerateSetupScript(t *testing.T) {
	manager := NewNginxManager()
	
	// 创建临时目录用于测试
	tempDir := t.TempDir()
	
	// 测试生成设置脚本
	err := manager.GenerateSetupScript(tempDir)
	if err != nil {
		t.Errorf("生成nginx设置脚本失败: %v", err)
	}
	
	// 验证脚本文件是否生成
	scriptFile := filepath.Join(tempDir, "setup-nginx.sh")
	if _, err := os.Stat(scriptFile); os.IsNotExist(err) {
		t.Errorf("nginx设置脚本应该被生成: %s", scriptFile)
	}
	
	// 验证脚本文件内容
	content, err := os.ReadFile(scriptFile)
	if err != nil {
		t.Errorf("读取设置脚本失败: %v", err)
	}
	
	contentStr := string(content)
	
	// 验证脚本包含必要的命令
	expectedItems := []string{
		"#!/bin/bash",
		"nginx -t",
		"systemctl reload nginx",
	}
	
	for _, item := range expectedItems {
		if !strings.Contains(contentStr, item) {
			t.Errorf("设置脚本应该包含: %s", item)
		}
	}
}

// TestGenerateSSLScript 测试生成SSL证书申请脚本
func TestGenerateSSLScript(t *testing.T) {
	manager := NewNginxManager()
	
	// 创建临时目录用于测试
	tempDir := t.TempDir()
	
	// 测试生成SSL脚本
	err := manager.GenerateSSLScript(tempDir)
	if err != nil {
		t.Errorf("生成SSL证书申请脚本失败: %v", err)
	}
	
	// 验证脚本文件是否生成
	scriptFile := filepath.Join(tempDir, "apply-ssl.sh")
	if _, err := os.Stat(scriptFile); os.IsNotExist(err) {
		t.Errorf("SSL证书申请脚本应该被生成: %s", scriptFile)
	}
	
	// 验证脚本文件内容
	content, err := os.ReadFile(scriptFile)
	if err != nil {
		t.Errorf("读取SSL脚本失败: %v", err)
	}
	
	contentStr := string(content)
	
	// 验证脚本包含必要的命令和逻辑
	expectedItems := []string{
		"#!/bin/bash",
		"certbot",
		"nginx -t",
		"systemctl reload nginx",
		"apt install",
		"yum install",
		"dnf install",
	}
	
	for _, item := range expectedItems {
		if !strings.Contains(contentStr, item) {
			t.Errorf("SSL脚本应该包含: %s", item)
		}
	}
}

// TestGenerateAll 测试生成所有nginx相关文件
func TestGenerateAll(t *testing.T) {
	manager := NewNginxManager()
	
	// 创建临时目录用于测试
	tempDir := t.TempDir()
	
	// 测试生成所有文件
	domain := "test.example.com"
	port := "8888"
	
	err := manager.GenerateAll(domain, tempDir, port)
	if err != nil {
		t.Errorf("生成所有nginx文件失败: %v", err)
	}
	
	// 验证所有文件都被生成
	expectedFiles := []string{
		domain,
		"setup-nginx.sh",
		"apply-ssl.sh",
	}
	
	for _, filename := range expectedFiles {
		filePath := filepath.Join(tempDir, filename)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("文件应该被生成: %s", filePath)
		}
	}
}

// TestGenerateAllWithEmptyPort 测试使用空端口生成所有文件
func TestGenerateAllWithEmptyPort(t *testing.T) {
	manager := NewNginxManager()
	
	// 创建临时目录用于测试
	tempDir := t.TempDir()
	
	// 测试生成所有文件（空端口）
	domain := "test.example.com"
	port := ""
	
	err := manager.GenerateAll(domain, tempDir, port)
	if err != nil {
		t.Errorf("生成所有nginx文件失败: %v", err)
	}
	
	// 验证配置文件被生成
	configFile := filepath.Join(tempDir, domain)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		t.Errorf("nginx配置文件应该被生成: %s", configFile)
	}
}

// TestGenerateConfigWithSpecialCharacters 测试包含特殊字符的域名
func TestGenerateConfigWithSpecialCharacters(t *testing.T) {
	manager := NewNginxManager()
	
	// 创建临时目录用于测试
	tempDir := t.TempDir()
	
	// 测试包含特殊字符的域名
	domain := "test-domain_with.dots-and-dashes.com"
	port := "8888"
	
	err := manager.GenerateConfig(domain, tempDir, port)
	if err != nil {
		t.Errorf("生成nginx配置失败: %v", err)
	}
	
	// 验证配置文件是否生成
	configFile := filepath.Join(tempDir, domain)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		t.Errorf("nginx配置文件应该被生成: %s", configFile)
	}
}

// TestGenerateConfigWithInvalidPath 测试无效路径
func TestGenerateConfigWithInvalidPath(t *testing.T) {
	manager := NewNginxManager()
	
	// 测试无效路径
	domain := "test.example.com"
	invalidPath := "/invalid/path/that/does/not/exist"
	port := "8888"
	
	err := manager.GenerateConfig(domain, invalidPath, port)
	if err == nil {
		t.Error("使用无效路径应该返回错误")
	}
}

// TestGenerateSetupScriptWithInvalidPath 测试无效路径的设置脚本生成
func TestGenerateSetupScriptWithInvalidPath(t *testing.T) {
	manager := NewNginxManager()
	
	// 测试无效路径
	invalidPath := "/invalid/path/that/does/not/exist"
	
	err := manager.GenerateSetupScript(invalidPath)
	if err == nil {
		t.Error("使用无效路径应该返回错误")
	}
}

// TestGenerateSSLScriptWithInvalidPath 测试无效路径的SSL脚本生成
func TestGenerateSSLScriptWithInvalidPath(t *testing.T) {
	manager := NewNginxManager()
	
	// 测试无效路径
	invalidPath := "/invalid/path/that/does/not/exist"
	
	err := manager.GenerateSSLScript(invalidPath)
	if err == nil {
		t.Error("使用无效路径应该返回错误")
	}
}

// TestNginxManagerConcurrent 测试NginxManager的并发安全性
func TestNginxManagerConcurrent(t *testing.T) {
	manager := NewNginxManager()
	
	// 创建临时目录用于测试
	tempDir := t.TempDir()
	
	// 创建多个goroutine同时使用manager
	done := make(chan bool, 3)
	
	go func() {
		_ = manager.GenerateConfig("domain1.com", tempDir, "8888")
		done <- true
	}()
	
	go func() {
		_ = manager.GenerateSetupScript(tempDir)
		done <- true
	}()
	
	go func() {
		_ = manager.GenerateSSLScript(tempDir)
		done <- true
	}()
	
	// 等待所有goroutine完成
	for i := 0; i < 3; i++ {
		<-done
	}
	
	// 如果没有panic，测试通过
}

// TestNginxManagerWithLogger 测试NginxManager与Logger的集成
func TestNginxManagerWithLogger(t *testing.T) {
	manager := NewNginxManager()
	logger := NewLogger()
	
	// 创建临时目录用于测试
	tempDir := t.TempDir()
	
	// 测试生成配置
	err := manager.GenerateConfig("test.example.com", tempDir, "8888")
	if err != nil {
		logger.Error("生成nginx配置失败: " + err.Error())
	} else {
		logger.Success("nginx配置生成成功")
	}
	
	// 如果代码执行到这里没有panic，测试通过
}

// TestNginxConfigContent 测试nginx配置内容
func TestNginxConfigContent(t *testing.T) {
	manager := NewNginxManager()
	
	// 创建临时目录用于测试
	tempDir := t.TempDir()
	
	// 生成配置
	domain := "test.example.com"
	port := "8888"
	
	err := manager.GenerateConfig(domain, tempDir, port)
	if err != nil {
		t.Fatalf("生成nginx配置失败: %v", err)
	}
	
	// 读取配置文件
	configFile := filepath.Join(tempDir, domain)
	content, err := os.ReadFile(configFile)
	if err != nil {
		t.Fatalf("读取配置文件失败: %v", err)
	}
	
	contentStr := string(content)
	
	// 验证配置文件的完整结构
	expectedSections := []string{
		"server {",
		"listen 80;",
		"server_name " + domain + ";",
		"location / {",
		"proxy_pass http://localhost:" + port + ";",
		"proxy_set_header Host $host;",
		"proxy_set_header X-Real-IP $remote_addr;",
		"proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;",
		"proxy_set_header X-Forwarded-Proto $scheme;",
		"proxy_http_version 1.1;",
		"proxy_set_header Upgrade $http_upgrade;",
		"proxy_set_header Connection \"upgrade\";",
		"proxy_connect_timeout 60s;",
		"proxy_send_timeout 60s;",
		"proxy_read_timeout 60s;",
		"}",
		"access_log /var/log/nginx/" + domain + ".access.log;",
		"error_log /var/log/nginx/" + domain + ".error.log;",
		"}",
	}
	
	for _, section := range expectedSections {
		if !strings.Contains(contentStr, section) {
			t.Errorf("配置文件应该包含: %s", section)
		}
	}
} 