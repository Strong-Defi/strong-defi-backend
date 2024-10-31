package log

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

// Logger 定义日志接口
type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	// 从 context 中获取字段
	WithContext(ctx context.Context) Logger
	WithFields(fields map[string]interface{}) Logger
	With(args ...interface{}) Logger
}

// ZapLogger 实现 Logger 接口
type ZapLogger struct {
	sugarLogger *zap.SugaredLogger
}

// Config 日志配置
type Config struct {
	Level          string                 // 日志级别 debug, info, warn, error, fatal
	Filename       string                 // 日志文件路径
	MaxSize        int                    // 每个日志文件最大尺寸(MB)
	MaxBackups     int                    // 保留的旧日志文件最大数量
	MaxAge         int                    // 保留的旧日志文件最大天数
	Compress       bool                   // 是否压缩旧日志文件
	Development    bool                   // 开发模式（更详细的堆栈跟踪）
	Encoding       string                 // 日志编码格式 (json/console)
	OutputPaths    []string               // 输出路径，支持 stdout/stderr/file
	InitialFields  map[string]interface{} // 初始字段
	DisableCaller  bool                   // 禁用调用者信息
	DisableTime    bool                   // 禁用时间戳
	TimeFormat     string                 // 时间格式
	SamplingConfig *SamplingConfig        // 采样配置
}

// SamplingConfig 采样配置
type SamplingConfig struct {
	Initial    int // 初始采样速率
	Thereafter int // 之后的采样速率
}

// contextKey 定义上下文键类型
type contextKey string

const (
	// TraceIDKey 定义跟踪ID的键
	TraceIDKey contextKey = "trace_id"
	// UserIDKey 定义用户ID的键
	UserIDKey contextKey = "user_id"
)

// Option 定义配置选项
type Option func(*Config)

// WithLevel 设置日志级别
func WithLevel(level string) Option {
	return func(c *Config) {
		c.Level = level
	}
}

// WithFilename 设置日志文件
func WithFilename(filename string) Option {
	return func(c *Config) {
		c.Filename = filename
	}
}

// WithDevelopment 设置开发模式
func WithDevelopment(enabled bool) Option {
	return func(c *Config) {
		c.Development = enabled
	}
}

// WithEncoding 设置日志编码格式
func WithEncoding(encoding string) Option {
	return func(c *Config) {
		c.Encoding = encoding
	}
}

// WithOutputPaths 设置输出路径
func WithOutputPaths(paths ...string) Option {
	return func(c *Config) {
		c.OutputPaths = paths
	}
}

// WithInitialFields 设置初始字段
func WithInitialFields(fields map[string]interface{}) Option {
	return func(c *Config) {
		c.InitialFields = fields
	}
}

// WithDisableCaller 设置禁用调用者信息
func WithDisableCaller(disable bool) Option {
	return func(c *Config) {
		c.DisableCaller = disable
	}
}

// WithDisableTime 设置禁用时间戳
func WithDisableTime(disable bool) Option {
	return func(c *Config) {
		c.DisableTime = disable
	}
}

// WithTimeFormat 设置时间格式
func WithTimeFormat(format string) Option {
	return func(c *Config) {
		c.TimeFormat = format
	}
}

// WithSampling 设置采样配置
func WithSampling(initial, thereafter int) Option {
	return func(c *Config) {
		c.SamplingConfig = &SamplingConfig{
			Initial:    initial,
			Thereafter: thereafter,
		}
	}
}

// WithMaxSize 设置每个日志文件最大尺寸
func WithMaxSize(size int) Option {
	return func(c *Config) {
		c.MaxSize = size
	}
}

// WithMaxBackups 设置保留的旧日志文件最大数量
func WithMaxBackups(count int) Option {
	return func(c *Config) {
		c.MaxBackups = count
	}
}

// WithMaxAge 设置保留的旧日志文件最大天数
func WithMaxAge(days int) Option {
	return func(c *Config) {
		c.MaxAge = days
	}
}

// WithCompress 设置是否压缩旧日志文件
func WithCompress(compress bool) Option {
	return func(c *Config) {
		c.Compress = compress
	}
}

// NewLogger 支持选项模式
func NewLogger(config *Config, opts ...Option) (Logger, error) {
	if config == nil {
		config = &Config{
			Level:       "info",
			MaxSize:     100,
			MaxBackups:  3,
			MaxAge:      28,
			Compress:    true,
			Encoding:    "json",
			OutputPaths: []string{"stdout"},
			TimeFormat:  "2006-01-02 15:04:05.000",
			Development: false,
		}
	}
	for _, opt := range opts {
		opt(config)
	}

	// 设置日志级别
	level := zapcore.InfoLevel
	switch config.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	}

	// 配置编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}

	// 应用时间格式配置
	if config.DisableTime {
		encoderConfig.TimeKey = ""
	} else {
		encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(config.TimeFormat))
		}
	}

	// 创建 zap 配置
	zapConfig := zap.Config{
		Level:         zap.NewAtomicLevelAt(level),
		Development:   config.Development,
		Encoding:      config.Encoding,
		EncoderConfig: encoderConfig,
		OutputPaths:   config.OutputPaths,
		InitialFields: config.InitialFields,
		DisableCaller: config.DisableCaller,
	}

	// 应用采样配置
	if config.SamplingConfig != nil {
		zapConfig.Sampling = &zap.SamplingConfig{
			Initial:    config.SamplingConfig.Initial,
			Thereafter: config.SamplingConfig.Thereafter,
		}
	}

	// 创建核心配置
	var core zapcore.Core
	if config.Filename != "" {
		// 文件输出
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.Filename,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			Compress:   config.Compress,
		})
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			writer,
			level,
		)
	} else {
		// 控制台输出
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			level,
		)
	}

	// 创建logger
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger := logger.Sugar()

	return &ZapLogger{
		sugarLogger: sugarLogger,
	}, nil
}

// 实现Logger接口的方法
func (l *ZapLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *ZapLogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *ZapLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *ZapLogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *ZapLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *ZapLogger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *ZapLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *ZapLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *ZapLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *ZapLogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}

// WithContext 从上下文创建新的带字段的 logger
func (l *ZapLogger) WithContext(ctx context.Context) Logger {
	if ctx == nil {
		return l
	}

	fields := extractFields(ctx)
	if len(fields) == 0 {
		return l
	}

	return &ZapLogger{
		sugarLogger: l.sugarLogger.With(fields...),
	}
}

// ContextField 定义上下文字段提取器
type ContextField struct {
	Key   string
	Value func(ctx context.Context) interface{}
}

var contextFields = []ContextField{
	{Key: "trace_id", Value: func(ctx context.Context) interface{} { return ctx.Value(TraceIDKey) }},
	{Key: "user_id", Value: func(ctx context.Context) interface{} { return ctx.Value(UserIDKey) }},
}

// RegisterContextField 注册新的上下文字段
func RegisterContextField(key string, extractor func(ctx context.Context) interface{}) {
	contextFields = append(contextFields, ContextField{Key: key, Value: extractor})
}

// 更新 extractFields 方法
func extractFields(ctx context.Context) []interface{} {
	var fields []interface{}
	for _, cf := range contextFields {
		if v := cf.Value(ctx); v != nil {
			fields = append(fields, cf.Key, v)
		}
	}
	return fields
}

// ZapLogger 实现新方法
func (l *ZapLogger) WithFields(fields map[string]interface{}) Logger {
	args := make([]interface{}, 0, len(fields)*2)
	for k, v := range fields {
		args = append(args, k, v)
	}
	return &ZapLogger{
		sugarLogger: l.sugarLogger.With(args...),
	}
}

func (l *ZapLogger) With(args ...interface{}) Logger {
	return &ZapLogger{
		sugarLogger: l.sugarLogger.With(args...),
	}
}
