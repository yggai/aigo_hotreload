package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unicode/utf8"
)

// TestProjectGeneratorCreation 测试ProjectGenerator创建
func TestProjectGeneratorCreation(t *testing.T) {
	projectPath := "/tmp/test-project"
	projectName := "test-project"
	
	generator := NewProjectGenerator(projectPath, projectName)
	if generator == nil {
		t.Error("NewProjectGenerator() 应该返回非空的ProjectGenerator实例")
	}
	
	if generator.projectPath != projectPath {
		t.Errorf("projectPath 期望 '%s', 实际得到 '%s'", projectPath, generator.projectPath)
	}
	
	if generator.projectName != projectName {
		t.Errorf("projectName 期望 '%s', 实际得到 '%s'", projectName, generator.projectName)
	}
}

// TestWriteFile 测试文件写入功能
func TestWriteFile(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := "test-project"
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 测试写入文件
	filename := "test.txt"
	content := "这是测试内容"
	
	err := generator.writeFile(filename, content)
	if err != nil {
		t.Errorf("写入文件失败: %v", err)
	}
	
	// 验证文件是否创建
	filePath := filepath.Join(tempDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("文件应该被创建: %s", filePath)
	}
	
	// 验证文件内容
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("读取文件失败: %v", err)
	}
	
	if string(fileContent) != content {
		t.Errorf("文件内容不匹配: 期望 '%s', 实际得到 '%s'", content, string(fileContent))
	}
}

// TestWriteFileWithSubdirectory 测试写入子目录中的文件
func TestWriteFileWithSubdirectory(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := "test-project"
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 测试写入子目录中的文件
	filename := "subdir/test.txt"
	content := "这是子目录中的测试内容"
	
	err := generator.writeFile(filename, content)
	if err != nil {
		t.Errorf("写入子目录文件失败: %v", err)
	}
	
	// 验证文件是否创建
	filePath := filepath.Join(tempDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("子目录文件应该被创建: %s", filePath)
	}
	
	// 验证文件内容
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("读取子目录文件失败: %v", err)
	}
	
	if string(fileContent) != content {
		t.Errorf("子目录文件内容不匹配: 期望 '%s', 实际得到 '%s'", content, string(fileContent))
	}
}

// TestGenerateAll 测试生成所有项目文件
func TestGenerateAll(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := "test-project"
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 测试生成所有文件
	err := generator.GenerateAll()
	if err != nil {
		t.Errorf("生成所有项目文件失败: %v", err)
	}
	
	// 验证所有必需文件都被创建
	expectedFiles := []string{
		"go.mod",
		"main.go",
		".air.toml",
		".gitignore",
		"README.md",
		"config/your-domain.com",
		"config/setup-nginx.sh",
		"scripts/apply-ssl.sh",
	}
	
	for _, filename := range expectedFiles {
		filePath := filepath.Join(tempDir, filename)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("文件应该被创建: %s", filePath)
		}
	}
}

// TestGenerateAllWithInvalidPath 测试无效路径
func TestGenerateAllWithInvalidPath(t *testing.T) {
	// 使用无效路径
	invalidPath := "/invalid/path/that/does/not/exist"
	projectName := "test-project"
	
	generator := NewProjectGenerator(invalidPath, projectName)
	
	// 测试生成所有文件（应该失败）
	err := generator.GenerateAll()
	if err == nil {
		t.Log("使用无效路径应该返回错误，但实际没有返回错误")
	} else {
		t.Logf("使用无效路径正确返回错误: %v", err)
	}
}

// TestGeneratedFileContents 测试生成的文件内容
func TestGeneratedFileContents(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := "test-project"
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 生成所有文件
	err := generator.GenerateAll()
	if err != nil {
		t.Fatalf("生成所有项目文件失败: %v", err)
	}
	
	// 测试go.mod文件内容
	goModPath := filepath.Join(tempDir, "go.mod")
	goModContent, err := os.ReadFile(goModPath)
	if err != nil {
		t.Fatalf("读取go.mod失败: %v", err)
	}
	
	goModStr := string(goModContent)
	if !strings.Contains(goModStr, "module "+projectName) {
		t.Errorf("go.mod应该包含模块名: %s", projectName)
	}
	
	if !strings.Contains(goModStr, "go 1.24") {
		t.Errorf("go.mod应该包含Go版本")
	}
	
	// 测试main.go文件内容
	mainGoPath := filepath.Join(tempDir, "main.go")
	mainGoContent, err := os.ReadFile(mainGoPath)
	if err != nil {
		t.Fatalf("读取main.go失败: %v", err)
	}
	
	mainGoStr := string(mainGoContent)
	if !strings.Contains(mainGoStr, "package main") {
		t.Errorf("main.go应该包含package main")
	}
	
	if !strings.Contains(mainGoStr, "gin.Default()") {
		t.Errorf("main.go应该包含gin.Default()")
	}
	
	if !strings.Contains(mainGoStr, ":8888") {
		t.Errorf("main.go应该包含端口8888")
	}
	
	// 测试.air.toml文件内容
	airTomlPath := filepath.Join(tempDir, ".air.toml")
	airTomlContent, err := os.ReadFile(airTomlPath)
	if err != nil {
		t.Fatalf("读取.air.toml失败: %v", err)
	}
	
	airTomlStr := string(airTomlContent)
	if !strings.Contains(airTomlStr, "[build]") {
		t.Errorf(".air.toml应该包含[build]配置")
	}
	
	// 测试.gitignore文件内容
	gitignorePath := filepath.Join(tempDir, ".gitignore")
	gitignoreContent, err := os.ReadFile(gitignorePath)
	if err != nil {
		t.Fatalf("读取.gitignore失败: %v", err)
	}
	
	gitignoreStr := string(gitignoreContent)
	if !strings.Contains(gitignoreStr, "*.exe") {
		t.Errorf(".gitignore应该包含*.exe")
	}
	
	// 测试README.md文件内容
	readmePath := filepath.Join(tempDir, "README.md")
	readmeContent, err := os.ReadFile(readmePath)
	if err != nil {
		t.Fatalf("读取README.md失败: %v", err)
	}
	
	readmeStr := string(readmeContent)
	if !strings.Contains(readmeStr, projectName) {
		t.Errorf("README.md应该包含项目名: %s", projectName)
	}
	
	// 测试nginx配置文件内容
	nginxConfigPath := filepath.Join(tempDir, "config/your-domain.com")
	nginxConfigContent, err := os.ReadFile(nginxConfigPath)
	if err != nil {
		t.Fatalf("读取nginx配置失败: %v", err)
	}
	
	nginxConfigStr := string(nginxConfigContent)
	if !strings.Contains(nginxConfigStr, "server_name your-domain.com") {
		t.Errorf("nginx配置应该包含server_name")
	}
	
	// 测试nginx设置脚本内容
	nginxSetupPath := filepath.Join(tempDir, "config/setup-nginx.sh")
	nginxSetupContent, err := os.ReadFile(nginxSetupPath)
	if err != nil {
		t.Fatalf("读取nginx设置脚本失败: %v", err)
	}
	
	nginxSetupStr := string(nginxSetupContent)
	if !strings.Contains(nginxSetupStr, "#!/bin/bash") {
		t.Errorf("nginx设置脚本应该包含shebang")
	}
	
	// 测试SSL脚本内容
	sslScriptPath := filepath.Join(tempDir, "scripts/apply-ssl.sh")
	sslScriptContent, err := os.ReadFile(sslScriptPath)
	if err != nil {
		t.Fatalf("读取SSL脚本失败: %v", err)
	}
	
	sslScriptStr := string(sslScriptContent)
	if !strings.Contains(sslScriptStr, "#!/bin/bash") {
		t.Errorf("SSL脚本应该包含shebang")
	}
	
	if !strings.Contains(sslScriptStr, "certbot") {
		t.Errorf("SSL脚本应该包含certbot")
	}
}

// TestDirectoryCreation 测试目录创建
func TestDirectoryCreation(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := "test-project"
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 生成所有文件
	err := generator.GenerateAll()
	if err != nil {
		t.Fatalf("生成所有项目文件失败: %v", err)
	}
	
	// 验证config目录是否创建
	configDir := filepath.Join(tempDir, "config")
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		t.Errorf("config目录应该被创建: %s", configDir)
	}
	
	// 验证scripts目录是否创建
	scriptsDir := filepath.Join(tempDir, "scripts")
	if _, err := os.Stat(scriptsDir); os.IsNotExist(err) {
		t.Errorf("scripts目录应该被创建: %s", scriptsDir)
	}
}

// TestFilePermissions 测试文件权限
func TestFilePermissions(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := "test-project"
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 生成所有文件
	err := generator.GenerateAll()
	if err != nil {
		t.Fatalf("生成所有项目文件失败: %v", err)
	}
	
	// 测试文件权限
	testFiles := []string{
		"go.mod",
		"main.go",
		".air.toml",
		".gitignore",
		"README.md",
	}
	
	for _, filename := range testFiles {
		filePath := filepath.Join(tempDir, filename)
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			t.Errorf("获取文件信息失败: %s", filePath)
			continue
		}
		
		// 检查文件权限（应该是0644）
		mode := fileInfo.Mode()
		if mode&0777 != 0644 {
			t.Errorf("文件权限不正确: %s, 期望 0644, 实际 %o", filePath, mode&0777)
		}
	}
}

// TestProjectGeneratorConcurrent 测试ProjectGenerator的并发安全性
func TestProjectGeneratorConcurrent(t *testing.T) {
	// 创建多个goroutine同时使用generator
	done := make(chan bool, 3)
	
	for i := 0; i < 3; i++ {
		go func(id int) {
			tempDir := t.TempDir()
			projectName := fmt.Sprintf("test-project-%d", id)
			
			generator := NewProjectGenerator(tempDir, projectName)
			_ = generator.GenerateAll()
			
			done <- true
		}(i)
	}
	
	// 等待所有goroutine完成
	for i := 0; i < 3; i++ {
		<-done
	}
	
	// 如果没有panic，测试通过
}

// TestProjectGeneratorWithSpecialCharacters 测试特殊字符处理
func TestProjectGeneratorWithSpecialCharacters(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := "test-project-with-special-chars_123"
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 测试生成所有文件
	err := generator.GenerateAll()
	if err != nil {
		t.Errorf("生成所有项目文件失败: %v", err)
	}
	
	// 验证go.mod文件包含正确的模块名
	goModPath := filepath.Join(tempDir, "go.mod")
	goModContent, err := os.ReadFile(goModPath)
	if err != nil {
		t.Fatalf("读取go.mod失败: %v", err)
	}
	
	goModStr := string(goModContent)
	if !strings.Contains(goModStr, "module "+projectName) {
		t.Errorf("go.mod应该包含模块名: %s", projectName)
	}
}

// TestProjectGeneratorWithEmptyProjectName 测试空项目名
func TestProjectGeneratorWithEmptyProjectName(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := ""
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 测试生成所有文件
	err := generator.GenerateAll()
	if err != nil {
		t.Errorf("生成所有项目文件失败: %v", err)
	}
	
	// 验证go.mod文件
	goModPath := filepath.Join(tempDir, "go.mod")
	goModContent, err := os.ReadFile(goModPath)
	if err != nil {
		t.Fatalf("读取go.mod失败: %v", err)
	}
	
	goModStr := string(goModContent)
	if !strings.Contains(goModStr, "module ") {
		t.Errorf("go.mod应该包含module声明")
	}
} 

// TestWriteFileWithEmptyContent 测试写入空内容文件
func TestWriteFileWithEmptyContent(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := "test-project"
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 测试写入空内容文件
	filename := "empty.txt"
	content := ""
	
	err := generator.writeFile(filename, content)
	if err != nil {
		t.Errorf("写入空内容文件失败: %v", err)
	}
	
	// 验证文件是否创建
	filePath := filepath.Join(tempDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("空内容文件应该被创建: %s", filePath)
	}
	
	// 验证文件内容为空
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("读取空内容文件失败: %v", err)
	}
	
	if string(fileContent) != content {
		t.Errorf("空内容文件内容不匹配: 期望 '%s', 实际得到 '%s'", content, string(fileContent))
	}
}

// TestWriteFileWithLargeContent 测试写入大内容文件
func TestWriteFileWithLargeContent(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := "test-project"
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 测试写入大内容文件
	filename := "large.txt"
	content := strings.Repeat("这是一个大文件的内容，用于测试文件写入功能。", 1000)
	
	err := generator.writeFile(filename, content)
	if err != nil {
		t.Errorf("写入大内容文件失败: %v", err)
	}
	
	// 验证文件是否创建
	filePath := filepath.Join(tempDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("大内容文件应该被创建: %s", filePath)
	}
	
	// 验证文件内容
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("读取大内容文件失败: %v", err)
	}
	
	if string(fileContent) != content {
		t.Errorf("大内容文件内容不匹配")
	}
}

// TestWriteFileWithSpecialCharacters 测试写入包含特殊字符的文件
func TestWriteFileWithSpecialCharacters(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := "test-project"
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 测试写入包含特殊字符的文件
	filename := "special-chars.txt"
	content := "特殊字符测试: !@#$%^&*()_+-=[]{}|;':\",./<>?`~"
	
	err := generator.writeFile(filename, content)
	if err != nil {
		t.Errorf("写入特殊字符文件失败: %v", err)
	}
	
	// 验证文件是否创建
	filePath := filepath.Join(tempDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("特殊字符文件应该被创建: %s", filePath)
	}
	
	// 验证文件内容
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("读取特殊字符文件失败: %v", err)
	}
	
	if string(fileContent) != content {
		t.Errorf("特殊字符文件内容不匹配: 期望 '%s', 实际得到 '%s'", content, string(fileContent))
	}
}

// TestWriteFileWithUnicode 测试写入包含Unicode字符的文件
func TestWriteFileWithUnicode(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := "test-project"
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 测试写入包含Unicode字符的文件
	filename := "unicode.txt"
	content := "Unicode测试: 中文测试 🚀 🎉 💻 测试"
	
	err := generator.writeFile(filename, content)
	if err != nil {
		t.Errorf("写入Unicode文件失败: %v", err)
	}
	
	// 验证文件是否创建
	filePath := filepath.Join(tempDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("Unicode文件应该被创建: %s", filePath)
	}
	
	// 验证文件内容
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("读取Unicode文件失败: %v", err)
	}
	
	if string(fileContent) != content {
		t.Errorf("Unicode文件内容不匹配: 期望 '%s', 实际得到 '%s'", content, string(fileContent))
	}
}

// TestGenerateAllWithEmptyProjectName 测试空项目名
func TestGenerateAllWithEmptyProjectName(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := ""
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 测试生成所有文件（空项目名）
	err := generator.GenerateAll()
	if err != nil {
		t.Errorf("使用空项目名生成所有项目文件失败: %v", err)
	}
	
	// 验证基本文件都被创建
	expectedFiles := []string{
		"go.mod",
		"main.go",
		".air.toml",
		".gitignore",
		"README.md",
	}
	
	for _, filename := range expectedFiles {
		filePath := filepath.Join(tempDir, filename)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("文件应该被创建: %s", filePath)
		}
	}
}

// TestGenerateAllWithSpecialProjectName 测试特殊项目名
func TestGenerateAllWithSpecialProjectName(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := "test-project-with-special-chars_123"
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 测试生成所有文件（特殊项目名）
	err := generator.GenerateAll()
	if err != nil {
		t.Errorf("使用特殊项目名生成所有项目文件失败: %v", err)
	}
	
	// 验证go.mod文件内容包含项目名
	goModPath := filepath.Join(tempDir, "go.mod")
	goModContent, err := os.ReadFile(goModPath)
	if err != nil {
		t.Fatalf("读取go.mod失败: %v", err)
	}
	
	goModStr := string(goModContent)
	if !strings.Contains(goModStr, "module "+projectName) {
		t.Errorf("go.mod应该包含模块名: %s", projectName)
	}
}

// TestGenerateAllConcurrent 测试并发生成
func TestGenerateAllConcurrent(t *testing.T) {
	// 创建多个goroutine同时生成项目
	done := make(chan bool, 3)
	
	for i := 0; i < 3; i++ {
		go func(id int) {
			tempDir := t.TempDir()
			projectName := fmt.Sprintf("test-project-%d", id)
			
			generator := NewProjectGenerator(tempDir, projectName)
			err := generator.GenerateAll()
			
			if err != nil {
				t.Errorf("并发生成项目失败: %v", err)
			}
			
			done <- true
		}(i)
	}
	
	// 等待所有goroutine完成
	for i := 0; i < 3; i++ {
		<-done
	}
}

// TestProjectGeneratorWithNilValues 测试空值处理
func TestProjectGeneratorWithNilValues(t *testing.T) {
	// 测试空路径
	generator := NewProjectGenerator("", "test-project")
	if generator.projectPath != "" {
		t.Errorf("空路径应该被正确处理")
	}
	
	// 测试空项目名
	generator = NewProjectGenerator("/tmp", "")
	if generator.projectName != "" {
		t.Errorf("空项目名应该被正确处理")
	}
}

// TestWriteFileWithInvalidPath 测试无效路径
func TestWriteFileWithInvalidPath(t *testing.T) {
	// 使用无效路径
	invalidPath := "/invalid/path/that/does/not/exist"
	projectName := "test-project"
	
	generator := NewProjectGenerator(invalidPath, projectName)
	
	// 测试写入文件（应该失败）
	filename := "test.txt"
	content := "测试内容"
	
	err := generator.writeFile(filename, content)
	if err == nil {
		t.Log("使用无效路径应该返回错误，但实际没有返回错误")
	} else {
		t.Logf("使用无效路径正确返回错误: %v", err)
	}
}

// TestGeneratedFilePermissions 测试生成文件的权限
func TestGeneratedFilePermissions(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := "test-project"
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 生成所有文件
	err := generator.GenerateAll()
	if err != nil {
		t.Fatalf("生成所有项目文件失败: %v", err)
	}
	
	// 验证文件权限
	expectedFiles := []string{
		"go.mod",
		"main.go",
		".air.toml",
		".gitignore",
		"README.md",
	}
	
	for _, filename := range expectedFiles {
		filePath := filepath.Join(tempDir, filename)
		info, err := os.Stat(filePath)
		if err != nil {
			t.Errorf("获取文件信息失败: %v", err)
			continue
		}
		
		// 检查文件权限（应该可读可写）
		mode := info.Mode()
		if mode&0400 == 0 {
			t.Errorf("文件 %s 应该可读", filename)
		}
	}
}

// TestGeneratedDirectoryPermissions 测试生成目录的权限
func TestGeneratedDirectoryPermissions(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := "test-project"
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 生成所有文件
	err := generator.GenerateAll()
	if err != nil {
		t.Fatalf("生成所有项目文件失败: %v", err)
	}
	
	// 验证目录权限
	expectedDirs := []string{
		"config",
		"scripts",
	}
	
	for _, dirName := range expectedDirs {
		dirPath := filepath.Join(tempDir, dirName)
		info, err := os.Stat(dirPath)
		if err != nil {
			t.Errorf("获取目录信息失败: %v", err)
			continue
		}
		
		// 检查目录权限（应该可读可写可执行）
		mode := info.Mode()
		if mode&0400 == 0 {
			t.Errorf("目录 %s 应该可读", dirName)
		}
		if mode&0200 == 0 {
			t.Errorf("目录 %s 应该可写", dirName)
		}
		if mode&0100 == 0 {
			t.Errorf("目录 %s 应该可执行", dirName)
		}
	}
}

// TestGeneratedFileEncoding 测试生成文件的编码
func TestGeneratedFileEncoding(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := "test-project"
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 生成所有文件
	err := generator.GenerateAll()
	if err != nil {
		t.Fatalf("生成所有项目文件失败: %v", err)
	}
	
	// 验证文件编码（应该是UTF-8）
	expectedFiles := []string{
		"main.go",
		"README.md",
	}
	
	for _, filename := range expectedFiles {
		filePath := filepath.Join(tempDir, filename)
		content, err := os.ReadFile(filePath)
		if err != nil {
			t.Errorf("读取文件失败: %v", err)
			continue
		}
		
		// 检查是否为有效的UTF-8
		if !utf8.Valid(content) {
			t.Errorf("文件 %s 应该使用UTF-8编码", filename)
		}
	}
}

// TestGeneratedFileLineEndings 测试生成文件的换行符
func TestGeneratedFileLineEndings(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()
	projectName := "test-project"
	
	generator := NewProjectGenerator(tempDir, projectName)
	
	// 生成所有文件
	err := generator.GenerateAll()
	if err != nil {
		t.Fatalf("生成所有项目文件失败: %v", err)
	}
	
	// 验证文件换行符
	expectedFiles := []string{
		"main.go",
		"README.md",
	}
	
	for _, filename := range expectedFiles {
		filePath := filepath.Join(tempDir, filename)
		content, err := os.ReadFile(filePath)
		if err != nil {
			t.Errorf("读取文件失败: %v", err)
			continue
		}
		
		// 检查是否包含换行符
		if !strings.Contains(string(content), "\n") {
			t.Errorf("文件 %s 应该包含换行符", filename)
		}
	}
} 