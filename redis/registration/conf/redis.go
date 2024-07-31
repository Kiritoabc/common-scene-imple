package conf

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

// RedisClient redis客户端
var RedisClient *Redis

// Redis redis客户端
type Redis struct {
	*redis.Client
}

// RedisConfig redis配置
type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// Init 初始化配置
func (c *RedisConfig) Init() (err error) {
	if len(c.Addr) == 0 {
		return errors.New("redis config is invalid")
	}
	RedisClient, err = NewClient(context.Background(), c)
	if err != nil {
		return
	}
	return nil
}

// NewClient 创建redis客户端
func NewClient(ctx context.Context, cfg *RedisConfig) (*Redis, error) {
	log.Info("redis client init")
	return &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr:     cfg.Addr,
			DB:       cfg.DB,
			Password: cfg.Password,
		}),
	}, nil
}
