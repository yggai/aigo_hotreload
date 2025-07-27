package generator

import (
	"fmt"
	"os"
	"path/filepath"

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

	return nil
}

// writeFile 写入文件
func (pg *ProjectGenerator) writeFile(filename, content string) error {
	filePath := filepath.Join(pg.projectPath, filename)
	return os.WriteFile(filePath, []byte(content), 0644)
}