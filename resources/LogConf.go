package resources

type Logging struct {
	IsStdout   bool   // 是否输出到控制台
	Filename   string //日志路径
	Prefix     string //日志前缀
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}
