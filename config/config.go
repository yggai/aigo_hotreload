package config

// Version CLI工具版本号
const Version = "1.0.0"

// DefaultPort 默认端口
const DefaultPort = "8888"

// Messages 消息文本
var Messages = struct {
	ToolName        string
	ToolDescription string
	UsageHeader     string
	Commands        struct {
		Create  string
		Version string
		Help    string
	}
	Example         string
	Errors          struct {
		NoProjectName   string
		EmptyName       string
		GetCwd          string
		DirExists       string
		CreateDir       string
		GenerateFiles   string
		UnknownCommand  string
	}
	Success         struct {
		Creating        string
		Created         string
		NextSteps       string
		Commands        struct {
			Cd      string
			Tidy    string
			Air     string
			Access  string
		}
	}
}{
	ToolName:        "aigo_hotreload",
	ToolDescription: "aigo_hotreload - Go热重载项目脚手架工具",
	UsageHeader:     "用法:",
	Commands: struct {
		Create  string
		Version string
		Help    string
	}{
		Create:  "  aigo_hotreload create <project-name>  创建新的热重载项目",
		Version: "  aigo_hotreload version               显示版本信息",
		Help:    "  aigo_hotreload help                  显示帮助信息",
	},
	Example: "示例:\n  aigo_hotreload create my-api",
	Errors: struct {
		NoProjectName   string
		EmptyName       string
		GetCwd          string
		DirExists       string
		CreateDir       string
		GenerateFiles   string
		UnknownCommand  string
	}{
		NoProjectName:  "错误: 请提供项目名称\n用法: aigo_hotreload create <project-name>",
		EmptyName:      "错误: 项目名称不能为空",
		GetCwd:         "错误: 无法获取当前目录: %v",
		DirExists:      "错误: 目录 %s 已存在",
		CreateDir:      "错误: 无法创建目录 %s: %v",
		GenerateFiles:  "错误: 生成项目文件失败: %v",
		UnknownCommand: "未知命令: %s",
	},
	Success: struct {
		Creating        string
		Created         string
		NextSteps       string
		Commands        struct {
			Cd      string
			Tidy    string
			Air     string
			Access  string
		}
	}{
		Creating:  "正在创建项目: %s",
		Created:   "✅ 项目 %s 创建成功!",
		NextSteps: "下一步:",
		Commands: struct {
			Cd      string
			Tidy    string
			Air     string
			Access  string
		}{
			Cd:     "  cd %s",
			Tidy:   "  go mod tidy",
			Air:    "  air",
			Access: "然后访问 http://localhost:8888 查看效果",
		},
	},
}