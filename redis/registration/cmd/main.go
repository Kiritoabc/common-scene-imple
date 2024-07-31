package main

import (
	"context"
	"github.com/kiritoabc/common-scene-imple/redis/registration/cmd/server"
)

// filePath 配置文件路径
const filePath = "conf/application.yaml"

func main() {
	server.Run(context.Background(), filePath)
}
