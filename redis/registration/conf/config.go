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
	App         App         `yaml:"app"`
	Log         Log         `yaml:"log"`
	RedisConfig RedisConfig `yaml:"redis"`
}

// App 应用配置
type App struct {
	Port       string `yaml:"port"`
	ServerName string `yaml:"server_name"`
}

// Log 日志配置
type Log struct {
	LogDir string        `yaml:"log_dir"`
	Level  string        `yaml:"level"`
	MaxAge time.Duration `yaml:"max_age"`
}
