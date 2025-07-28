package config

import (
	"strings"
	"testing"
)

// TestConstants 测试常量定义
func TestConstants(t *testing.T) {
	// 测试服务器相关常量
	if DefaultPort != "8888" {
		t.Errorf("DefaultPort 期望 '8888', 实际得到 '%s'", DefaultPort)
	}

	if DefaultHost != "localhost" {
		t.Errorf("DefaultHost 期望 'localhost', 实际得到 '%s'", DefaultHost)
	}

	// 测试URL相关常量
	expectedBaseURL := "http://localhost:8888"
	if BaseURL != expectedBaseURL {
		t.Errorf("BaseURL 期望 '%s', 实际得到 '%s'", expectedBaseURL, BaseURL)
	}

	expectedHealthURL := "http://localhost:8888/health"
	if HealthURL != expectedHealthURL {
		t.Errorf("HealthURL 期望 '%s', 实际得到 '%s'", expectedHealthURL, HealthURL)
	}

	expectedAPIBaseURL := "http://localhost:8888/api/v1"
	if APIBaseURL != expectedAPIBaseURL {
		t.Errorf("APIBaseURL 期望 '%s', 实际得到 '%s'", expectedAPIBaseURL, APIBaseURL)
	}

	expectedUsersAPIURL := "http://localhost:8888/api/v1/users"
	if UsersAPIURL != expectedUsersAPIURL {
		t.Errorf("UsersAPIURL 期望 '%s', 实际得到 '%s'", expectedUsersAPIURL, UsersAPIURL)
	}

	// 测试Air相关常量
	if AirInstallCmd != "go install github.com/air-verse/air@latest" {
		t.Errorf("AirInstallCmd 期望 'go install github.com/air-verse/air@latest', 实际得到 '%s'", AirInstallCmd)
	}

	if AirCommand != "air" {
		t.Errorf("AirCommand 期望 'air', 实际得到 '%s'", AirCommand)
	}

	// 测试文件权限常量
	if DirPermission != 0755 {
		t.Errorf("DirPermission 期望 0755, 实际得到 %o", DirPermission)
	}

	if FilePermission != 0644 {
		t.Errorf("FilePermission 期望 0644, 实际得到 %o", FilePermission)
	}

	// 测试状态消息常量
	if StatusSuccess != "success" {
		t.Errorf("StatusSuccess 期望 'success', 实际得到 '%s'", StatusSuccess)
	}

	if StatusHealthy != "healthy" {
		t.Errorf("StatusHealthy 期望 'healthy', 实际得到 '%s'", StatusHealthy)
	}

	// 测试时间格式常量
	if TimeFormat != "2006-01-02 15:04:05" {
		t.Errorf("TimeFormat 期望 '2006-01-02 15:04:05', 实际得到 '%s'", TimeFormat)
	}

	if DateFormat != "2006-01-02" {
		t.Errorf("DateFormat 期望 '2006-01-02', 实际得到 '%s'", DateFormat)
	}
}

// TestConfigMessages 测试配置消息
func TestConfigMessages(t *testing.T) {
	// 测试成功消息
	if Messages.Success.Created == "" {
		t.Error("Messages.Success.Created 不能为空")
	}

	if Messages.Success.Creating == "" {
		t.Error("Messages.Success.Creating 不能为空")
	}

	if Messages.Success.NextSteps == "" {
		t.Error("Messages.Success.NextSteps 不能为空")
	}

	// 测试错误消息
	if Messages.Errors.NoProjectName == "" {
		t.Error("Messages.Errors.NoProjectName 不能为空")
	}

	if Messages.Errors.EmptyName == "" {
		t.Error("Messages.Errors.EmptyName 不能为空")
	}

	if Messages.Errors.UnknownCommand == "" {
		t.Error("Messages.Errors.UnknownCommand 不能为空")
	}

	// 测试成功消息
	if Messages.Success.NextSteps != "下一步:" {
		t.Errorf("Messages.Success.NextSteps 期望 '下一步:', 实际得到 '%s'", Messages.Success.NextSteps)
	}

	if Messages.Success.Created == "" {
		t.Error("Messages.Success.Created 不能为空")
	}

	if Messages.Success.Creating == "" {
		t.Error("Messages.Success.Creating 不能为空")
	}

	// 测试命令消息
	if !strings.Contains(Messages.Commands.Create, "创建新的热重载项目") {
		t.Errorf("Messages.Commands.Create 应该包含 '创建新的热重载项目', 实际得到 '%s'", Messages.Commands.Create)
	}

	if !strings.Contains(Messages.Commands.Help, "显示帮助信息") {
		t.Errorf("Messages.Commands.Help 应该包含 '显示帮助信息', 实际得到 '%s'", Messages.Commands.Help)
	}

	if !strings.Contains(Messages.Commands.Version, "显示版本信息") {
		t.Errorf("Messages.Commands.Version 应该包含 '显示版本信息', 实际得到 '%s'", Messages.Commands.Version)
	}

	if !strings.Contains(Messages.Commands.Nginx, "生成nginx配置") {
		t.Errorf("Messages.Commands.Nginx 应该包含 '生成nginx配置', 实际得到 '%s'", Messages.Commands.Nginx)
	}

	// 测试访问消息
	if Messages.Success.Commands.Access == "" {
		t.Error("Messages.Success.Commands.Access 不能为空")
	}

	// 测试部署消息
	if Messages.Success.Deployment.NginxSetup == "" {
		t.Error("Messages.Success.Deployment.NginxSetup 不能为空")
	}

	if Messages.Success.Deployment.SSLSetup == "" {
		t.Error("Messages.Success.Deployment.SSLSetup 不能为空")
	}

	if Messages.Success.Deployment.DomainAccess == "" {
		t.Error("Messages.Success.Deployment.DomainAccess 不能为空")
	}
}

// TestConfigStructure 测试配置结构完整性
func TestConfigStructure(t *testing.T) {
	// 测试Messages结构不为空
	if Messages.Success.Created == "" {
		t.Error("Messages.Success.Created 不能为空")
	}

	if Messages.Errors.NoProjectName == "" {
		t.Error("Messages.Errors.NoProjectName 不能为空")
	}

	if Messages.Success.NextSteps == "" {
		t.Error("Messages.Success.NextSteps 不能为空")
	}

	if Messages.Commands.Create == "" {
		t.Error("Messages.Commands.Create 不能为空")
	}

	if Messages.Success.Commands.Access == "" {
		t.Error("Messages.Success.Commands.Access 不能为空")
	}

	if Messages.Success.Deployment.NginxSetup == "" {
		t.Error("Messages.Success.Deployment.NginxSetup 不能为空")
	}
}

// TestURLConstruction 测试URL构建逻辑
func TestURLConstruction(t *testing.T) {
	// 验证BaseURL构建逻辑
	expectedBaseURL := "http://" + DefaultHost + ":" + DefaultPort
	if BaseURL != expectedBaseURL {
		t.Errorf("BaseURL构建错误: 期望 '%s', 实际得到 '%s'", expectedBaseURL, BaseURL)
	}

	// 验证HealthURL构建逻辑
	expectedHealthURL := BaseURL + "/health"
	if HealthURL != expectedHealthURL {
		t.Errorf("HealthURL构建错误: 期望 '%s', 实际得到 '%s'", expectedHealthURL, HealthURL)
	}

	// 验证APIBaseURL构建逻辑
	expectedAPIBaseURL := BaseURL + "/api/v1"
	if APIBaseURL != expectedAPIBaseURL {
		t.Errorf("APIBaseURL构建错误: 期望 '%s', 实际得到 '%s'", expectedAPIBaseURL, APIBaseURL)
	}

	// 验证UsersAPIURL构建逻辑
	expectedUsersAPIURL := APIBaseURL + "/users"
	if UsersAPIURL != expectedUsersAPIURL {
		t.Errorf("UsersAPIURL构建错误: 期望 '%s', 实际得到 '%s'", expectedUsersAPIURL, UsersAPIURL)
	}
} 