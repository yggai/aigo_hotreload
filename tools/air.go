package tools

import (
	"os"
	"os/exec"

	"github.com/yggai/aigo_hotreload/config"
)

// AirManager Air 工具管理器
type AirManager struct {
	logger *Logger
}

// NewAirManager 创建新的 Air 管理器
func NewAirManager() *AirManager {
	return &AirManager{
		logger: NewLogger(),
	}
}

// InstallAir 检查并安装 air 工具
func (m *AirManager) InstallAir() {
	// 检查 air 是否已安装
	if _, err := exec.LookPath(config.AirCommand); err == nil {
		m.logger.Success("Air 已安装")
		return
	}

	m.logger.Info("正在安装 Air 热重载工具...")
	cmd := exec.Command("go", "install", config.AirInstallCmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		m.logger.Warning("Air 安装失败: %v", err)
		m.logger.Info("请手动安装: %s", config.AirInstallCmd)
		return
	}

	m.logger.Success("Air 安装成功")
} 