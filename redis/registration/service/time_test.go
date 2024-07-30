package service

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	// 获取当前的年份
	year := now.Year()
	// 获取当前日期是今年的第几天
	dayOfYear := now.YearDay()
	month := now.Month().String()
	fmt.Println("当前年份：==", year)
	fmt.Println("当前月份：==", month)
	fmt.Println("当前日期是今年的第几天：==", dayOfYear)
}

func TestName(t *testing.T) {
	now := time.Now()
	// 获取当前月的天数
	days := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location()).Add(-24 * time.Hour).Day()
	t.Log(days)
	offset := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).YearDay()
	t.Log(offset)
}
