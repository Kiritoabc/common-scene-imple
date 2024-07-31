package server

import (
	"context"
	"github.com/kiritoabc/common-scene-imple/redis/registration/conf"
	"github.com/kiritoabc/common-scene-imple/redis/registration/service"
	log "github.com/sirupsen/logrus"
)
import "github.com/gin-gonic/gin"

// userSvc 用户服务（简单模拟）
var userSvc = &service.UserSvc{}

// Run 启动服务
func Run(ctx context.Context, filePath string) {
	installConfigOrDie(filePath)
	if err := installPlugins(); err != nil {
		log.Fatalln(err)
	}
	engine := gin.Default()
	// 签到
	engine.POST("/register", userSvc.Register)
	// 获取一年签到的天数
	engine.GET("/cumulative_days", userSvc.GetCumulativeDays)
	// 获取当月的签到情况
	engine.GET("/sign_of_month", userSvc.GetSignOfMonth)
	err := engine.Run(conf.Config.Config.App.Port)
	if err != nil {
		return
	}
}
