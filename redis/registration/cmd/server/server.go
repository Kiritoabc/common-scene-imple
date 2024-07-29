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

	engine := gin.Default()
	// 签到
	engine.POST("/register", userSvc.Register)

	log.Fatalln(engine.Run(conf.Config.Config.App.Port))
}
