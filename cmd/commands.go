package cmd

import (
	"os"

	"github.com/yggai/aigo_hotreload/config"
	"github.com/yggai/aigo_hotreload/project"
	"github.com/yggai/aigo_hotreload/tools"
)

// CommandHandler 命令处理器
type CommandHandler struct {
	projectManager *project.Manager
	logger         *tools.Logger
}

// NewCommandHandler 创建新的命令处理器
func NewCommandHandler() *CommandHandler {
	return &CommandHandler{
		projectManager: project.NewManager(),
		logger:         tools.NewLogger(),
	}
}

// HandleCommands 处理命令行命令
func (h *CommandHandler) HandleCommands() {
	if len(os.Args) < 2 {
		h.printUsage()
		return
	}

	command := os.Args[1]
	switch command {
	case "create":
		h.handleCreate()
	case "nginx":
		h.handleNginx()
	case "version":
		h.handleVersion()
	case "help":
		h.printUsage()
	default:
		h.logger.Error(config.Messages.Errors.UnknownCommand, command)
		h.printUsage()
	}
}

// handleCreate 处理创建项目命令
func (h *CommandHandler) handleCreate() {
	if len(os.Args) < 3 {
		h.logger.Error(config.Messages.Errors.NoProjectName)
		return
	}
	projectName := os.Args[2]
	h.projectManager.CreateProject(projectName)
}

// handleVersion 处理版本命令
func (h *CommandHandler) handleVersion() {
	h.logger.Info("aigo_hotreload version %s", config.Version)
}

// handleNginx 处理nginx配置命令
func (h *CommandHandler) handleNginx() {
	if len(os.Args) < 4 {
		h.logger.Error("用法: aigo_hotreload nginx <domain> <project-path> [port]")
		return
	}
	
	domain := os.Args[2]
	projectPath := os.Args[3]
	port := ""
	
	if len(os.Args) >= 5 {
		port = os.Args[4]
	}
	
	nginxManager := tools.NewNginxManager()
	if err := nginxManager.GenerateAll(domain, projectPath, port); err != nil {
		h.logger.Error("生成nginx配置失败: %v", err)
		return
	}
	
	h.logger.Success("nginx配置生成完成")
	h.logger.Info("下一步:")
	h.logger.Info("1. 编辑配置文件: vim %s/config/%s", projectPath, domain)
	h.logger.Info("2. 运行配置脚本: %s/config/setup-nginx.sh %s", projectPath, domain)
	h.logger.Info("3. 申请SSL证书: %s/scripts/apply-ssl.sh %s", projectPath, domain)
}

// printUsage 打印使用说明
func (h *CommandHandler) printUsage() {
	h.logger.Println(config.Messages.ToolDescription)
	h.logger.PrintEmpty()
	h.logger.Println(config.Messages.UsageHeader)
	h.logger.Println(config.Messages.Commands.Create)
	h.logger.Println(config.Messages.Commands.Nginx)
	h.logger.Println(config.Messages.Commands.Version)
	h.logger.Println(config.Messages.Commands.Help)
	h.logger.PrintEmpty()
	h.logger.Println(config.Messages.Example)
} 