package log

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	// 测试默认配置
	logger, err := NewLogger(nil)
	if err != nil {
		t.Fatalf("NewLogger() error = %v", err)
	}
	if logger == nil {
		t.Fatal("NewLogger() returned nil logger")
	}

	// 测试文件输出
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "test.log")

	fileLogger, err := NewLogger(nil,
		WithLevel("debug"),
		WithFilename(logFile),
	)
	if err != nil {
		t.Fatalf("NewLogger() with file error = %v", err)
	}

	fileLogger.Info("test file logging")

	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Error("Log file was not created")
	}
}

func TestLogLevels(t *testing.T) {
	logger, _ := NewLogger(nil)

	// 测试各个日志级别
	logger.Debug("debug message")
	logger.Debugf("debug %s", "formatted")
	logger.Info("info message")
	logger.Infof("info %s", "formatted")
	logger.Warn("warn message")
	logger.Warnf("warn %s", "formatted")
	logger.Error("error message")
	logger.Errorf("error %s", "formatted")

	// Fatal 级别不测试，因为会导致程序退出
}

func TestWithContext(t *testing.T) {
	logger, _ := NewLogger(nil)

	// 创建带有跟踪ID的上下文
	ctx := context.WithValue(context.Background(), TraceIDKey, "test-trace-id")
	ctx = context.WithValue(ctx, UserIDKey, "test-user-id")

	ctxLogger := logger.WithContext(ctx)
	ctxLogger.Info("log with context")

	// 测试空上下文
	nilCtxLogger := logger.WithContext(nil)
	if nilCtxLogger != logger {
		t.Error("WithContext(nil) should return original logger")
	}
}

func TestWithFields(t *testing.T) {
	logger, _ := NewLogger(nil)

	fields := map[string]interface{}{
		"key1": "value1",
		"key2": 123,
	}

	fieldLogger := logger.WithFields(fields)
	fieldLogger.Info("log with fields")
}

func TestWith(t *testing.T) {
	logger, _ := NewLogger(nil)

	withLogger := logger.With(
		"string", "value",
		"number", 42,
		"bool", true,
	)
	withLogger.Info("log with additional fields")
}

func TestRegisterContextField(t *testing.T) {
	// 注册新的上下文字段
	RegisterContextField("test_field", func(ctx context.Context) interface{} {
		return ctx.Value("test_key")
	})

	logger, _ := NewLogger(nil)
	ctx := context.WithValue(context.Background(), "test_key", "test_value")

	ctxLogger := logger.WithContext(ctx)
	ctxLogger.Info("log with custom context field")
}

func TestLoggerChaining(t *testing.T) {
	logger, _ := NewLogger(nil)

	// 测试方法链式调用
	ctx := context.WithValue(context.Background(), TraceIDKey, "trace-id")
	logger.
		WithContext(ctx).
		WithFields(map[string]interface{}{"key": "value"}).
		With("extra", "field").
		Info("chained logging")
}

func TestLoggerWithConfig(t *testing.T) {
	// 创建临时配置文件
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "logger.yml")
	logPath := filepath.Join(tmpDir, "logs/app.log")

	// 确保日志目录存在
	if err := os.MkdirAll(filepath.Join(tmpDir, "logs"), 0755); err != nil {
		t.Fatal(err)
	}

	// 创建测试配置
	configContent := []byte(`
logger:
  level: "debug"
  development: true
  encoding: "json"
  output_paths: 
    - "stdout"
    - "` + logPath + `"
  initial_fields:
    app: "test-service"
    env: "test"
  disable_caller: false
  disable_time: false
  time_format: "2006-01-02 15:04:05.000"
  sampling:
    initial: 100
    thereafter: 10
  file_config:
    max_size: 100
    max_backups: 3
    max_age: 28
    compress: true
`)

	if err := os.WriteFile(configPath, configContent, 0644); err != nil {
		t.Fatal(err)
	}

	// 从配置创建 logger
	logger, err := NewLoggerFromConfig(configPath)
	if err != nil {
		t.Fatalf("Failed to create logger from config: %v", err)
	}

	// 测试日志输出
	testCases := []struct {
		name    string
		logFunc func()
		message string
		fields  map[string]interface{}
	}{
		{
			name: "basic info log",
			logFunc: func() {
				logger.Info("test info message")
			},
			message: "test info message",
		},
		{
			name: "log with fields",
			logFunc: func() {
				logger.WithFields(map[string]interface{}{
					"test_key": "test_value",
				}).Info("test with fields")
			},
			message: "test with fields",
			fields: map[string]interface{}{
				"test_key": "test_value",
			},
		},
		{
			name: "log with context",
			logFunc: func() {
				ctx := context.WithValue(context.Background(), TraceIDKey, "test-trace-id")
				logger.WithContext(ctx).Info("test with context")
			},
			message: "test with context",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.logFunc()

			// 验证日志文件是否存在
			if _, err := os.Stat(logPath); os.IsNotExist(err) {
				t.Error("Log file was not created")
			}

			// 读取日志文件内容进行验证
			content, err := os.ReadFile(logPath)
			if err != nil {
				t.Fatal(err)
			}

			if !strings.Contains(string(content), tc.message) {
				t.Errorf("Log file does not contain expected message: %s", tc.message)
			}

			// 验证初始字段
			if !strings.Contains(string(content), "test-service") {
				t.Error("Log file does not contain initial field 'app'")
			}

			// 验证额外字段
			if tc.fields != nil {
				for k, v := range tc.fields {
					if !strings.Contains(string(content), fmt.Sprintf(`"%s":"%v"`, k, v)) {
						t.Errorf("Log file does not contain field %s=%v", k, v)
					}
				}
			}
		})
	}
}

func TestInvalidConfig(t *testing.T) {
	// 测试无效的配置文件路径
	_, err := NewLoggerFromConfig("non_existent_config.yml")
	if err == nil {
		t.Error("Expected error for non-existent config file")
	}

	// 测试无效的配置内容
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "invalid_logger.yml")

	invalidConfig := []byte(`
logger:
  level: "invalid_level"
`)

	if err := os.WriteFile(configPath, invalidConfig, 0644); err != nil {
		t.Fatal(err)
	}

	_, err = NewLoggerFromConfig(configPath)
	if err == nil {
		t.Error("Expected error for invalid log level")
	}
}
