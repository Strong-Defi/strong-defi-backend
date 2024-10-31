package log

import (
	"os"

	"gopkg.in/yaml.v3"
)

type LoggerConfig struct {
	Logger struct {
		Level         string                 `yaml:"level"`
		Development   bool                   `yaml:"development"`
		Encoding      string                 `yaml:"encoding"`
		OutputPaths   []string               `yaml:"output_paths"`
		InitialFields map[string]interface{} `yaml:"initial_fields"`
		DisableCaller bool                   `yaml:"disable_caller"`
		DisableTime   bool                   `yaml:"disable_time"`
		TimeFormat    string                 `yaml:"time_format"`
		Sampling      struct {
			Initial    int `yaml:"initial"`
			Thereafter int `yaml:"thereafter"`
		} `yaml:"sampling"`
		FileConfig struct {
			MaxSize    int  `yaml:"max_size"`
			MaxBackups int  `yaml:"max_backups"`
			MaxAge     int  `yaml:"max_age"`
			Compress   bool `yaml:"compress"`
		} `yaml:"file_config"`
	} `yaml:"logger"`
}

func NewLoggerFromConfig(configPath string) (Logger, error) {
	config := &LoggerConfig{}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}

	opts := []Option{
		WithLevel(config.Logger.Level),
		WithDevelopment(config.Logger.Development),
		WithEncoding(config.Logger.Encoding),
		WithOutputPaths(config.Logger.OutputPaths...),
		WithInitialFields(config.Logger.InitialFields),
		WithDisableCaller(config.Logger.DisableCaller),
		WithDisableTime(config.Logger.DisableTime),
		WithTimeFormat(config.Logger.TimeFormat),
		WithSampling(config.Logger.Sampling.Initial, config.Logger.Sampling.Thereafter),
		WithMaxSize(config.Logger.FileConfig.MaxSize),
		WithMaxBackups(config.Logger.FileConfig.MaxBackups),
		WithMaxAge(config.Logger.FileConfig.MaxAge),
		WithCompress(config.Logger.FileConfig.Compress),
	}

	return NewLogger(nil, opts...)
}
