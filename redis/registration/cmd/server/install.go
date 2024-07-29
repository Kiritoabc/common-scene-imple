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
	vip := viper.New()
	vip.AddConfigPath(filePath)      //设置读取的文件路径
	vip.SetConfigName("application") //设置读取的文件名
	vip.SetConfigType("yaml")        //设置文件的类型
	//尝试进行配置读取
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := vip.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	conf.Config = &conf.Provider{
		Config: cfg,
	}
}
