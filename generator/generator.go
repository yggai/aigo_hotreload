package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yggai/aigo_hotreload/config"
	"github.com/yggai/aigo_hotreload/templates"
)

// ProjectGenerator 项目生成器
type ProjectGenerator struct {
	projectPath string
	projectName string
}

// NewProjectGenerator 创建新的项目生成器
func NewProjectGenerator(projectPath, projectName string) *ProjectGenerator {
	return &ProjectGenerator{
		projectPath: projectPath,
		projectName: projectName,
	}
}

// GenerateAll 生成所有项目文件
func (pg *ProjectGenerator) GenerateAll() error {
	files := []struct {
		filename string
		content  string
	}{
		{"go.mod", fmt.Sprintf(templates.GoModTemplate, pg.projectName)},
		{"main.go", templates.MainGoTemplate},
		{".air.toml", templates.AirTomlTemplate},
		{".gitignore", templates.GitignoreTemplate},
		{"README.md", fmt.Sprintf(templates.ReadmeTemplate, pg.projectName, pg.projectName)},
	}

	for _, file := range files {
		if err := pg.writeFile(file.filename, file.content); err != nil {
			return fmt.Errorf("生成文件 %s 失败: %v", file.filename, err)
		}
	}

	// 创建config目录
	configDir := filepath.Join(pg.projectPath, "config")
	if err := os.MkdirAll(configDir, config.DirPermission); err != nil {
		return fmt.Errorf("创建config目录失败: %v", err)
	}

	// 生成nginx配置文件
	nginxFiles := []struct {
		filename string
		content  string
	}{
		{"config/your-domain.com", fmt.Sprintf(templates.NginxHTTPTemplate, "your-domain.com", config.DefaultPort, "your-domain.com", "your-domain.com")},
		{"config/setup-nginx.sh", templates.NginxSetupScriptTemplate},
	}

	for _, file := range nginxFiles {
		if err := pg.writeFile(file.filename, file.content); err != nil {
			return fmt.Errorf("生成文件 %s 失败: %v", file.filename, err)
		}
	}

	// 创建scripts目录
	scriptsDir := filepath.Join(pg.projectPath, "scripts")
	if err := os.MkdirAll(scriptsDir, config.DirPermission); err != nil {
		return fmt.Errorf("创建scripts目录失败: %v", err)
	}

	// 生成SSL证书申请脚本
	sslScriptFile := "scripts/apply-ssl.sh"
	if err := pg.writeFile(sslScriptFile, templates.CertbotScriptTemplate); err != nil {
		return fmt.Errorf("生成文件 %s 失败: %v", sslScriptFile, err)
	}

	return nil
}

// writeFile 写入文件
func (pg *ProjectGenerator) writeFile(filename, content string) error {
	filePath := filepath.Join(pg.projectPath, filename)
	
	// 确保目录存在
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, config.DirPermission); err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}
	
	return os.WriteFile(filePath, []byte(content), config.FilePermission)
}