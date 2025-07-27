package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yggai/aigo_hotreload/config"
	"github.com/yggai/aigo_hotreload/generator"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]
	switch command {
	case "create":
		if len(os.Args) < 3 {
			fmt.Println(config.Messages.Errors.NoProjectName)
			return
		}
		projectName := os.Args[2]
		createProject(projectName)
	case "version":
		fmt.Printf("aigo_hotreload version %s\n", config.Version)
	case "help":
		printUsage()
	default:
		fmt.Printf(config.Messages.Errors.UnknownCommand+"\n", command)
		printUsage()
	}
}

func printUsage() {
	fmt.Println(config.Messages.ToolDescription)
	fmt.Println("")
	fmt.Println(config.Messages.UsageHeader)
	fmt.Println(config.Messages.Commands.Create)
	fmt.Println(config.Messages.Commands.Version)
	fmt.Println(config.Messages.Commands.Help)
	fmt.Println("")
	fmt.Println(config.Messages.Example)
}

func createProject(projectName string) {
	// 检查项目名称是否有效
	if strings.TrimSpace(projectName) == "" {
		fmt.Println(config.Messages.Errors.EmptyName)
		return
	}

	// 获取当前工作目录
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf(config.Messages.Errors.GetCwd+"\n", err)
		return
	}

	projectPath := filepath.Join(cwd, projectName)

	// 检查目录是否已存在
	if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
		fmt.Printf(config.Messages.Errors.DirExists+"\n", projectPath)
		return
	}

	// 创建项目目录
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		fmt.Printf(config.Messages.Errors.CreateDir+"\n", projectPath, err)
		return
	}

	fmt.Printf(config.Messages.Success.Creating+"\n", projectName)

	// 生成项目文件
	gen := generator.NewProjectGenerator(projectPath, projectName)
	if err := gen.GenerateAll(); err != nil {
		fmt.Printf(config.Messages.Errors.GenerateFiles+"\n", err)
		return
	}

	fmt.Printf(config.Messages.Success.Created+"\n", projectName)
	fmt.Println("")
	fmt.Println(config.Messages.Success.NextSteps)
	fmt.Printf(config.Messages.Success.Commands.Cd+"\n", projectName)
	fmt.Println(config.Messages.Success.Commands.Tidy)
	fmt.Println(config.Messages.Success.Commands.Air)
	fmt.Println("")
	fmt.Println(config.Messages.Success.Commands.Access)
}
