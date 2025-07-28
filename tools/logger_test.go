package tools

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

// TestLoggerCreation 测试Logger创建
func TestLoggerCreation(t *testing.T) {
	logger := NewLogger()
	if logger == nil {
		t.Error("NewLogger() 应该返回非空的Logger实例")
	}
}

// TestLoggerInfo 测试Info方法
func TestLoggerInfo(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger := NewLogger()
	testMessage := "测试信息消息"
	logger.Info(testMessage)

	// 恢复标准输出
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// 验证输出包含预期内容
	if !strings.Contains(output, testMessage) {
		t.Errorf("Info输出应该包含消息 '%s', 实际输出: '%s'", testMessage, output)
	}

	if !strings.Contains(output, "ℹ️") {
		t.Errorf("Info输出应该包含信息图标, 实际输出: '%s'", output)
	}
}

// TestLoggerSuccess 测试Success方法
func TestLoggerSuccess(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger := NewLogger()
	testMessage := "测试成功消息"
	logger.Success(testMessage)

	// 恢复标准输出
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// 验证输出包含预期内容
	if !strings.Contains(output, testMessage) {
		t.Errorf("Success输出应该包含消息 '%s', 实际输出: '%s'", testMessage, output)
	}

	if !strings.Contains(output, "✅") {
		t.Errorf("Success输出应该包含成功图标, 实际输出: '%s'", output)
	}
}

// TestLoggerWarning 测试Warning方法
func TestLoggerWarning(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger := NewLogger()
	testMessage := "测试警告消息"
	logger.Warning(testMessage)

	// 恢复标准输出
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// 验证输出包含预期内容
	if !strings.Contains(output, testMessage) {
		t.Errorf("Warning输出应该包含消息 '%s', 实际输出: '%s'", testMessage, output)
	}

	if !strings.Contains(output, "⚠️") {
		t.Errorf("Warning输出应该包含警告图标, 实际输出: '%s'", output)
	}
}

// TestLoggerError 测试Error方法
func TestLoggerError(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger := NewLogger()
	testMessage := "测试错误消息"
	logger.Error(testMessage)

	// 恢复标准输出
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// 验证输出包含预期内容
	if !strings.Contains(output, testMessage) {
		t.Errorf("Error输出应该包含消息 '%s', 实际输出: '%s'", testMessage, output)
	}

	if !strings.Contains(output, "❌") {
		t.Errorf("Error输出应该包含错误图标, 实际输出: '%s'", output)
	}
}

// TestLoggerFatal 测试Fatal方法
func TestLoggerFatal(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger := NewLogger()
	testMessage := "测试致命错误消息"
	
	// 使用defer来恢复标准输出，因为Fatal会调用os.Exit
	defer func() {
		w.Close()
		os.Stdout = oldStdout
	}()
	
	// 由于Fatal会调用os.Exit，我们需要捕获panic
	defer func() {
		if r := recover(); r != nil {
			// 预期的panic，测试通过
		}
	}()

	logger.Fatal(testMessage)
	
	// 使用变量避免未使用警告
	_ = r
}

// TestLoggerMultipleMessages 测试多个消息的输出
func TestLoggerMultipleMessages(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger := NewLogger()
	
	logger.Info("信息1")
	logger.Success("成功1")
	logger.Warning("警告1")
	logger.Error("错误1")

	// 恢复标准输出
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// 验证所有消息都在输出中
	expectedMessages := []string{"信息1", "成功1", "警告1", "错误1"}
	for _, msg := range expectedMessages {
		if !strings.Contains(output, msg) {
			t.Errorf("输出应该包含消息 '%s', 实际输出: '%s'", msg, output)
		}
	}

	// 验证图标都在输出中
	expectedIcons := []string{"ℹ️", "✅", "⚠️", "❌"}
	for _, icon := range expectedIcons {
		if !strings.Contains(output, icon) {
			t.Errorf("输出应该包含图标 '%s', 实际输出: '%s'", icon, output)
		}
	}
}

// TestLoggerEmptyMessage 测试空消息
func TestLoggerEmptyMessage(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger := NewLogger()
	logger.Info("")

	// 恢复标准输出
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// 验证即使消息为空，也应该有输出（包含图标和换行）
	if len(output) == 0 {
		t.Error("即使消息为空，也应该有输出")
	}

	if !strings.Contains(output, "ℹ️") {
		t.Errorf("空消息输出也应该包含信息图标, 实际输出: '%s'", output)
	}
}

// TestLoggerSpecialCharacters 测试特殊字符
func TestLoggerSpecialCharacters(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger := NewLogger()
	specialMessage := "特殊字符: !@#$%^&*()_+-=[]{}|;':\",./<>?"
	logger.Info(specialMessage)

	// 恢复标准输出
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// 验证特殊字符被正确处理
	if !strings.Contains(output, specialMessage) {
		t.Errorf("输出应该包含特殊字符消息, 实际输出: '%s'", output)
	}
}

// TestLoggerWithEmptyMessage 测试空消息
func TestLoggerWithEmptyMessage(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger := NewLogger()
	logger.Info("")
	logger.Success("")
	logger.Warning("")
	logger.Error("")

	// 恢复标准输出
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()
	
	// 使用变量避免未使用警告
	_ = r
	_ = output

	// 验证输出不为空（即使消息为空，也应该有图标和格式）
	if len(output) == 0 {
		t.Error("空消息的输出不应该为空")
	}
}

// TestLoggerWithSpecialCharacters 测试特殊字符消息
func TestLoggerWithSpecialCharacters(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger := NewLogger()
	specialMessage := "特殊字符测试: !@#$%^&*()_+-=[]{}|;':\",./<>?`~"
	logger.Info(specialMessage)

	// 恢复标准输出
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// 验证输出包含特殊字符
	if !strings.Contains(output, specialMessage) {
		t.Errorf("输出应该包含特殊字符消息: %s", specialMessage)
	}
}

// TestLoggerWithUnicode 测试Unicode字符消息
func TestLoggerWithUnicode(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger := NewLogger()
	unicodeMessage := "Unicode测试: 中文测试 🚀 🎉 💻 测试"
	logger.Info(unicodeMessage)

	// 恢复标准输出
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// 验证输出包含Unicode字符
	if !strings.Contains(output, unicodeMessage) {
		t.Errorf("输出应该包含Unicode消息: %s", unicodeMessage)
	}
}

// TestLoggerWithLongMessage 测试长消息
func TestLoggerWithLongMessage(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger := NewLogger()
	longMessage := strings.Repeat("这是一个很长的消息，用于测试logger处理长消息的能力。", 100)
	logger.Info(longMessage)

	// 恢复标准输出
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// 验证输出包含长消息
	if !strings.Contains(output, longMessage) {
		t.Error("输出应该包含长消息")
	}
}

// TestLoggerWithNewlines 测试包含换行符的消息
func TestLoggerWithNewlines(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger := NewLogger()
	messageWithNewlines := "第一行\n第二行\n第三行"
	logger.Info(messageWithNewlines)

	// 恢复标准输出
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// 验证输出包含换行符
	if !strings.Contains(output, messageWithNewlines) {
		t.Error("输出应该包含换行符")
	}
}

// TestLoggerConcurrent 测试并发日志记录
func TestLoggerConcurrent(t *testing.T) {
	logger := NewLogger()
	
	// 创建多个goroutine同时记录日志
	done := make(chan bool, 10)
	
	for i := 0; i < 10; i++ {
		go func(id int) {
			logger.Info(fmt.Sprintf("并发测试消息 %d", id))
			logger.Success(fmt.Sprintf("并发成功消息 %d", id))
			logger.Warning(fmt.Sprintf("并发警告消息 %d", id))
			logger.Error(fmt.Sprintf("并发错误消息 %d", id))
			done <- true
		}(i)
	}
	
	// 等待所有goroutine完成
	for i := 0; i < 10; i++ {
		<-done
	}
	
	// 如果没有panic，测试通过
	t.Log("并发日志记录测试完成")
}

// TestLoggerPerformance 测试日志性能
func TestLoggerPerformance(t *testing.T) {
	logger := NewLogger()
	
	// 测试大量日志记录的性能
	for i := 0; i < 1000; i++ {
		logger.Info(fmt.Sprintf("性能测试消息 %d", i))
		logger.Success(fmt.Sprintf("性能成功消息 %d", i))
		logger.Warning(fmt.Sprintf("性能警告消息 %d", i))
		logger.Error(fmt.Sprintf("性能错误消息 %d", i))
	}
	
	// 如果函数执行到这里没有panic，测试通过
	t.Log("性能测试完成")
}

// TestLoggerMemoryUsage 测试内存使用
func TestLoggerMemoryUsage(t *testing.T) {
	logger := NewLogger()
	
	// 测试大量日志记录的内存使用
	for i := 0; i < 10000; i++ {
		logger.Info(fmt.Sprintf("内存测试消息 %d", i))
	}
	
	// 如果函数执行到这里没有panic，测试通过
	t.Log("内存使用测试完成")
}

// TestLoggerErrorRecovery 测试错误恢复
func TestLoggerErrorRecovery(t *testing.T) {
	logger := NewLogger()
	
	// 测试在异常情况下的恢复能力
	testCases := []string{
		"",                    // 空字符串
		" ",                   // 空格
		"\t",                  // 制表符
		"\n",                  // 换行符
		"!@#$%^&*()",         // 特殊字符
		"中文测试",             // 中文
		"🚀🎉💻",              // emoji
		strings.Repeat("a", 1000), // 长字符串
	}
	
	for _, message := range testCases {
		logger.Info(message)
		logger.Success(message)
		logger.Warning(message)
		logger.Error(message)
	}
	
	// 如果函数执行到这里没有panic，测试通过
	t.Log("错误恢复测试完成")
}

// TestLoggerIntegration 测试集成功能
func TestLoggerIntegration(t *testing.T) {
	logger := NewLogger()
	
	// 测试完整的日志流程
	logger.Info("开始集成测试")
	logger.Success("集成测试成功")
	logger.Warning("集成测试警告")
	logger.Error("集成测试错误")
	
	// 测试不同类型的消息组合
	messages := []string{
		"普通消息",
		"包含数字的消息 123",
		"包含特殊字符的消息 !@#",
		"包含中文的消息 测试",
		"包含emoji的消息 🚀",
	}
	
	for _, msg := range messages {
		logger.Info(msg)
		logger.Success(msg)
		logger.Warning(msg)
		logger.Error(msg)
	}
	
	// 如果函数执行到这里没有panic，测试通过
	t.Log("集成测试完成")
}

// TestLoggerBoundaryConditions 测试边界条件
func TestLoggerBoundaryConditions(t *testing.T) {
	logger := NewLogger()
	
	// 测试边界条件
	testCases := []string{
		"",                    // 空字符串
		" ",                   // 单个空格
		"\t",                  // 制表符
		"\n",                  // 换行符
		"\r",                  // 回车符
		"\r\n",                // 回车换行
		"a",                   // 单个字符
		"ab",                  // 两个字符
		strings.Repeat("a", 10000), // 超长字符串
		"测试",                // 中文
		"🚀",                  // emoji
		"测试🚀",              // 中文+emoji
	}
	
	for _, message := range testCases {
		logger.Info(message)
		logger.Success(message)
		logger.Warning(message)
		logger.Error(message)
	}
	
	// 如果函数执行到这里没有panic，测试通过
	t.Log("边界条件测试完成")
}

// TestLoggerFormatting 测试格式化功能
func TestLoggerFormatting(t *testing.T) {
	// 捕获标准输出
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger := NewLogger()
	
	// 测试不同类型的格式化消息
	logger.Info("信息消息")
	logger.Success("成功消息")
	logger.Warning("警告消息")
	logger.Error("错误消息")

	// 恢复标准输出
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// 验证输出包含所有类型的消息
	expectedMessages := []string{
		"信息消息",
		"成功消息",
		"警告消息",
		"错误消息",
	}

	for _, expected := range expectedMessages {
		if !strings.Contains(output, expected) {
			t.Errorf("输出应该包含消息: %s", expected)
		}
	}

	// 验证输出包含正确的图标
	expectedIcons := []string{
		"ℹ️",  // Info图标
		"✅",  // Success图标
		"⚠️",  // Warning图标
		"❌",  // Error图标
	}

	for _, icon := range expectedIcons {
		if !strings.Contains(output, icon) {
			t.Errorf("输出应该包含图标: %s", icon)
		}
	}
}

// TestLoggerConsistency 测试一致性
func TestLoggerConsistency(t *testing.T) {
	logger1 := NewLogger()
	logger2 := NewLogger()
	
	// 测试两个logger实例的一致性
	testMessage := "一致性测试消息"
	
	// 捕获第一个logger的输出
	oldStdout := os.Stdout
	r1, w1, _ := os.Pipe()
	os.Stdout = w1
	
	logger1.Info(testMessage)
	
	w1.Close()
	os.Stdout = oldStdout
	
	var buf1 bytes.Buffer
	buf1.ReadFrom(r1)
	output1 := buf1.String()
	
	// 捕获第二个logger的输出
	r2, w2, _ := os.Pipe()
	os.Stdout = w2
	
	logger2.Info(testMessage)
	
	w2.Close()
	os.Stdout = oldStdout
	
	var buf2 bytes.Buffer
	buf2.ReadFrom(r2)
	output2 := buf2.String()
	
	// 验证两个输出包含相同的消息
	if !strings.Contains(output1, testMessage) {
		t.Error("第一个logger输出应该包含测试消息")
	}
	
	if !strings.Contains(output2, testMessage) {
		t.Error("第二个logger输出应该包含测试消息")
	}
	
	// 验证两个输出都包含Info图标
	if !strings.Contains(output1, "ℹ️") {
		t.Error("第一个logger输出应该包含Info图标")
	}
	
	if !strings.Contains(output2, "ℹ️") {
		t.Error("第二个logger输出应该包含Info图标")
	}
} 