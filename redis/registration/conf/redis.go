package conf

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

var RedisClient = &Redis{
	Client: nil,
}

// Redis redis客户端
type Redis struct {
	*redis.Client
}

// RedisConfig redis配置
type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// Init 初始化配置
func (c *RedisConfig) Init() (err error) {
	if len(c.Addr) == 0 || len(c.Password) == 0 {
		return errors.New("redis config is invalid")
	}
	RedisClient, err = NewClient(context.Background(), &Config.Config.RedisConfig)
	if err != nil {
		return
	}
	return nil
}

// NewClient 创建redis客户端
func NewClient(ctx context.Context, cfg *RedisConfig) (*Redis, error) {
	log.Info("redis client init")
	return &Redis{
		redis.NewClient(&redis.Options{
			Addr:     cfg.Addr,
			Password: cfg.Password,
			DB:       cfg.DB,
		}),
	}, nil
}
