# 用户签到场景

##  key的设计

签到记录以年为单位，一个用户，对应一张位图（`Bitmap`），表示用户在一年内的签到情况。

- `key` 的设计：`user:sign:%d:%d`，第一个占位符表示年份，第二个占位符表示用户的编号。
- `bitmap` 值的设计：由于一年只有 **365** 或 **366** 天，因此我们只需要 `bitmap` 里面的前 **366** 位，即 **0-365** 位。





**如何获取当前年份？**

~~~go
	now := time.Now()
	// 获取当前的年份
	year := now.Year()
	fmt.Println("当前年份：==", year)
	// 获取当前日期是今年的第几天
	dayOfYear := now.YearDay()
	fmt.Println("当前日期是今年的第几天：==", dayOfYear)
~~~

