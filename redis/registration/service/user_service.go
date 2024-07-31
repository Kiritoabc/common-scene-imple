package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kiritoabc/common-scene-imple/redis/registration/conf"
	"github.com/kiritoabc/common-scene-imple/redis/registration/domain"
	"net/http"
	"time"
)

// UserSvc 用户服务（api+router+service）
type UserSvc struct{}

// Register 签到
func (s *UserSvc) Register(ctx *gin.Context) {
	user := &domain.User{}
	err := ctx.ShouldBindJSON(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "error",
		})
		return
	}
	now := time.Now()
	// 获取当前的年份
	year := now.Year()
	// 获取当前日期是今年的第几天
	dayOfYear := now.YearDay()
	// 签到 key: user:sign:年份:用户ID
	key := fmt.Sprintf("user:sign:%d:%d", year, user.ID)
	// setbit key offset value
	oldValue, err := conf.RedisClient.SetBit(ctx, key, int64(dayOfYear), 1).Result()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  "error",
		})
		return
	}
	if oldValue == 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "重复签到",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "签到成功",
	})
	return
}

// GetCumulativeDays 获取指定年份的累计签到天数
func (s *UserSvc) GetCumulativeDays(ctx *gin.Context) {
	// 获取用户信息
	userId := ctx.Query("user_id")
	// 获得时间
	now := time.Now()
	// 当前年份
	year := now.Year()
	// 当前天数的偏移量
	dayOfYear := now.YearDay()
	// 拼接key
	key := fmt.Sprintf("user:sign:%d:%s", year, userId)
	segmentSize := 63
	cumulativeDays := 0
	// bit操作
	bitOps := make([]any, 0)
	for i := 0; i < dayOfYear; i += segmentSize {
		size := segmentSize
		if i+segmentSize > dayOfYear {
			size = dayOfYear - i + 1
		}
		// GET, usize,#i
		// get,u25,190
		bitOps = append(bitOps, "GET", fmt.Sprintf("u%d", size), fmt.Sprintf("%d", i))
	}

	values, err := conf.RedisClient.BitField(ctx, key, bitOps...).Result()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  "error",
		})
		return
	}
	// 遍历
	for idx, value := range values {
		if value != 0 {
			size := segmentSize
			if (idx+1)*segmentSize > dayOfYear {
				size = dayOfYear % segmentSize
			}
			for j := 0; j < size; j++ {
				// 位运算判断结果
				if (value & (1 << (size - 1 - j))) != 0 {
					cumulativeDays++
				}
			}
		}
	}
	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": cumulativeDays,
	})
	return
}

// GetSignOfMonth 获取指定月份的签到情况
func (s *UserSvc) GetSignOfMonth(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	now := time.Now()
	year := now.Year()
	// 获取当前月的天数
	days := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location()).Add(-24 * time.Hour).Day() // 31
	// 获取本月初是今年的第几天
	offset := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).YearDay()
	key := fmt.Sprintf("user:sign:%d:%s", year, userId)
	typ := fmt.Sprintf("u%d", days)
	values, err := conf.RedisClient.BitField(ctx, key, "Get", typ, offset).Result()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  "error",
		})
		return
	}
	signInSlice := make([]bool, days)
	if len(values) == 0 {
		signInBits := values[0]
		for i := 0; i < days; i++ {
			signInSlice[i] = (signInBits & (1 << (days - 1 - i))) != 0
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": signInSlice,
	})
}
