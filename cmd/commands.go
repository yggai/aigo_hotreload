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

// printUsage 打印使用说明
func (h *CommandHandler) printUsage() {
	h.logger.Println(config.Messages.ToolDescription)
	h.logger.PrintEmpty()
	h.logger.Println(config.Messages.UsageHeader)
	h.logger.Println(config.Messages.Commands.Create)
	h.logger.Println(config.Messages.Commands.Version)
	h.logger.Println(config.Messages.Commands.Help)
	h.logger.PrintEmpty()
	h.logger.Println(config.Messages.Example)
} 