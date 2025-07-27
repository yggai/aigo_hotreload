package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yggai/aigo_hotreload/config"
	"github.com/yggai/aigo_hotreload/templates"
)

// NginxManager nginx配置管理器
type NginxManager struct {
	logger *Logger
}

// NewNginxManager 创建新的nginx管理器
func NewNginxManager() *NginxManager {
	return &NginxManager{
		logger: NewLogger(),
	}
}

// GenerateConfig 生成nginx配置文件
func (nm *NginxManager) GenerateConfig(domain, projectPath string, port string) error {
	if strings.TrimSpace(domain) == "" {
		return fmt.Errorf("域名不能为空")
	}

	if port == "" {
		port = config.DefaultPort
	}

	// 创建config目录
	configDir := filepath.Join(projectPath, "config")
	if err := os.MkdirAll(configDir, config.DirPermission); err != nil {
		return fmt.Errorf("创建config目录失败: %v", err)
	}

	// 生成nginx配置文件
	configFile := filepath.Join(configDir, domain)
	content := fmt.Sprintf(templates.NginxHTTPTemplate, domain, port, domain, domain)
	
	if err := os.WriteFile(configFile, []byte(content), config.FilePermission); err != nil {
		return fmt.Errorf("生成nginx配置文件失败: %v", err)
	}

	nm.logger.Success("nginx配置文件已生成: %s", configFile)
	return nil
}

// GenerateSetupScript 生成nginx配置脚本
func (nm *NginxManager) GenerateSetupScript(projectPath string) error {
	configDir := filepath.Join(projectPath, "config")
	setupScript := filepath.Join(configDir, "setup-nginx.sh")
	
	if err := os.WriteFile(setupScript, []byte(templates.NginxSetupScriptTemplate), config.FilePermission); err != nil {
		return fmt.Errorf("生成nginx配置脚本失败: %v", err)
	}

	// 添加执行权限
	if err := os.Chmod(setupScript, 0755); err != nil {
		return fmt.Errorf("设置脚本执行权限失败: %v", err)
	}

	nm.logger.Success("nginx配置脚本已生成: %s", setupScript)
	return nil
}

// GenerateSSLScript 生成SSL证书申请脚本
func (nm *NginxManager) GenerateSSLScript(projectPath string) error {
	scriptsDir := filepath.Join(projectPath, "scripts")
	if err := os.MkdirAll(scriptsDir, config.DirPermission); err != nil {
		return fmt.Errorf("创建scripts目录失败: %v", err)
	}

	sslScript := filepath.Join(scriptsDir, "apply-ssl.sh")
	
	if err := os.WriteFile(sslScript, []byte(templates.CertbotScriptTemplate), config.FilePermission); err != nil {
		return fmt.Errorf("生成SSL证书申请脚本失败: %v", err)
	}

	// 添加执行权限
	if err := os.Chmod(sslScript, 0755); err != nil {
		return fmt.Errorf("设置脚本执行权限失败: %v", err)
	}

	nm.logger.Success("SSL证书申请脚本已生成: %s", sslScript)
	return nil
}

// GenerateAll 生成所有nginx相关文件
func (nm *NginxManager) GenerateAll(domain, projectPath string, port string) error {
	if err := nm.GenerateConfig(domain, projectPath, port); err != nil {
		return err
	}

	if err := nm.GenerateSetupScript(projectPath); err != nil {
		return err
	}

	if err := nm.GenerateSSLScript(projectPath); err != nil {
		return err
	}

	nm.logger.Success("所有nginx相关文件已生成完成")
	return nil
} 