package main

import (
	"github.com/yggai/aigo_hotreload/cmd"
)

func main() {
	// 创建命令处理器并处理命令
	handler := cmd.NewCommandHandler()
	handler.HandleCommands()
}
