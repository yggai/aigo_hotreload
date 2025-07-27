package project

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/yggai/aigo_hotreload/config"
	"github.com/yggai/aigo_hotreload/generator"
	"github.com/yggai/aigo_hotreload/tools"
)

// Manager 项目管理器
type Manager struct {
	airManager *tools.AirManager
	logger     *tools.Logger
}

// NewManager 创建新的项目管理器
func NewManager() *Manager {
	return &Manager{
		airManager: tools.NewAirManager(),
		logger:     tools.NewLogger(),
	}
}

// CreateProject 创建新项目
func (m *Manager) CreateProject(projectName string) {
	// 检查项目名称是否有效
	if strings.TrimSpace(projectName) == "" {
		m.logger.Error(config.Messages.Errors.EmptyName)
		return
	}

	// 获取当前工作目录
	cwd, err := os.Getwd()
	if err != nil {
		m.logger.Error(config.Messages.Errors.GetCwd, err)
		return
	}

	projectPath := filepath.Join(cwd, projectName)

	// 检查目录是否已存在
	if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
		m.logger.Error(config.Messages.Errors.DirExists, projectPath)
		return
	}

	// 创建项目目录
	if err := os.MkdirAll(projectPath, config.DirPermission); err != nil {
		m.logger.Error(config.Messages.Errors.CreateDir, projectPath, err)
		return
	}

	m.logger.Info(config.Messages.Success.Creating, projectName)

	// 生成项目文件
	gen := generator.NewProjectGenerator(projectPath, projectName)
	if err := gen.GenerateAll(); err != nil {
		m.logger.Error(config.Messages.Errors.GenerateFiles, err)
		return
	}

	m.logger.Success(config.Messages.Success.Created, projectName)
	m.logger.PrintEmpty()

	// 检查并安装 air
	m.airManager.InstallAir()

	// 显示后续步骤
	m.showNextSteps(projectName)
}

// showNextSteps 显示项目创建后的后续步骤
func (m *Manager) showNextSteps(projectName string) {
	m.logger.Println(config.Messages.Success.NextSteps)
	m.logger.Info(config.Messages.Success.Commands.Cd, projectName)
	m.logger.Println(config.Messages.Success.Commands.Tidy)
	m.logger.Println(config.Messages.Success.Commands.Air)
	m.logger.PrintEmpty()
	m.logger.Println(config.Messages.Success.Commands.Access)
} 