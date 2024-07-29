package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kiritoabc/common-scene-imple/redis/registration/conf"
	"github.com/kiritoabc/common-scene-imple/redis/registration/domain"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

// UserSvc 用户服务（api+router+service）
type UserSvc struct{}

// Register 签到
func (s *UserSvc) Register(ctx *gin.Context) {
	user := &domain.User{}
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "error",
		})
	}
	now := time.Now()
	// 获取当前的年份
	year := now.Year()
	fmt.Println(year)
	// 获取当前日期是今年的第几天
	dayOfYear := now.YearDay()
	fmt.Println(dayOfYear)
	// 签到
	fmt.Sprintf("user:%s:%s", user.Name, "s")
	conf.RedisClient.SetBit(ctx, "user_"+user.Name, user.Age, 0)
}

// GetCumulativeDays 获取指定年份的累计签到天数
func GetCumulativeDays(ctx context.Context, rdb *redis.Client, userID int, year int, dayOfYear int) (int, error) {
	key := fmt.Sprintf("user:%d:%d", year, userID)
	segmentSize := 63
	cumulativeDays := 0
	bitOps := make([]any, 0)

	for i := 0; i < dayOfYear; i += segmentSize {
		size := segmentSize
		if i+segmentSize > dayOfYear {
			size = dayOfYear - i
		}

		bitOps = append(bitOps, "GET", fmt.Sprintf("u%d", size), fmt.Sprintf("#%d", i))
	}

	values, err := rdb.BitField(ctx, key, bitOps...).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to get bitfield: %w", err)
	}

	for idx, value := range values {
		if value != 0 {
			size := segmentSize
			if (idx+1)*segmentSize > dayOfYear {
				size = dayOfYear % segmentSize
			}
			for j := 0; j < size; j++ {
				if (value & (1 << (size - 1 - j))) != 0 {
					cumulativeDays++
				}
			}
		}
	}
	return cumulativeDays, nil
}

func GetSignOfMonth(ctx context.Context, rdb *redis.Client, userID, year, days, offset int) ([]bool, error) {
	typ := fmt.Sprintf("u%d", days)
	key := fmt.Sprintf("user:%d:%d", year, userID)

	s, err := rdb.BitField(ctx, key, "GET", typ, offset).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get bitfield: %w", err)
	}

	if len(s) != 0 {
		signInBits := s[0]
		signInSlice := make([]bool, days)
		for i := 0; i < days; i++ {
			signInSlice[i] = (signInBits & (1 << (days - 1 - i))) != 0
		}
		return signInSlice, nil
	} else {
		return nil, errors.New("no result returned from BITFIELD command")
	}
}
