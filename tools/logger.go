package tools

import (
	"fmt"
	"os"
)

// Logger 日志工具
type Logger struct{}

// NewLogger 创建新的日志工具
func NewLogger() *Logger {
	return &Logger{}
}

// Info 打印信息
func (l *Logger) Info(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

// Success 打印成功信息
func (l *Logger) Success(format string, args ...interface{}) {
	fmt.Printf("✅ "+format+"\n", args...)
}

// Warning 打印警告信息
func (l *Logger) Warning(format string, args ...interface{}) {
	fmt.Printf("⚠️  "+format+"\n", args...)
}

// Error 打印错误信息
func (l *Logger) Error(format string, args ...interface{}) {
	fmt.Printf("❌ "+format+"\n", args...)
}

// Print 直接打印
func (l *Logger) Print(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Println 打印并换行
func (l *Logger) Println(message string) {
	fmt.Println(message)
}

// PrintEmpty 打印空行
func (l *Logger) PrintEmpty() {
	fmt.Println("")
}

// Fatal 打印错误并退出
func (l *Logger) Fatal(format string, args ...interface{}) {
	fmt.Printf("💥 "+format+"\n", args...)
	os.Exit(1)
} 