package server

import (
	"github.com/kiritoabc/common-scene-imple/redis/registration/conf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// installConfigOrDie 安装配置文件
func installConfigOrDie(filePath string) {
	log.Info("Installing config file")
	cfg := &conf.Configuration{}
	viper.SetConfigFile(filePath)
	viper.SetConfigType("yaml") //设置文件的类型
	//尝试进行配置读取
	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("Failed to read config: %v", err)
		return // 同上
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Errorf("Failed to unmarshal config: %v", err)
		return // 同上
	}
	conf.Config = &conf.Provider{
		Config: cfg,
	}
}

func installPlugins() error {
	return conf.Config.Config.RedisConfig.Init()
}
