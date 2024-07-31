package conf

import "time"

// Provider 配置提供者
type Provider struct {
	Config *Configuration
}

// Config 配置实体（简单操作一下）
var Config = &Provider{}

// Configuration 配置文件
type Configuration struct {
	App         App         `mapstructure:"app"`
	Log         Log         `mapstructure:"log"`
	RedisConfig RedisConfig `mapstructure:"redis"`
}

// App 应用配置
type App struct {
	Port       string `mapstructure:"port"`
	ServerName string `mapstructure:"server_name"`
}

// Log 日志配置
type Log struct {
	LogDir string        `mapstructure:"log_dir"`
	Level  string        `mapstructure:"level"`
	MaxAge time.Duration `mapstructure:"max_age"`
}
