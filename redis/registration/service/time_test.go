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
