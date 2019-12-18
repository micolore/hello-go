package main

import (
	"fmt"
	"time"
)

// Duration 持续时间这个api 有点意思
func main() {

	//month_str := time.Month.String
	//fmt.Printf("month string: %v", month_str)
	year_day := time.Now().YearDay()
	fmt.Printf("now year day: %d\n", year_day)
	year := time.Now().Year()
	fmt.Printf("now year: %d\n", year)
	month := time.Now().Month()
	fmt.Printf("now month: %d\n", month)
	day := time.Now().Day()
	fmt.Printf("now day: %d\n", day)
	year_month_day := fmt.Sprintf("%d-%d-%d", year, month, day)
	yearmonthday := fmt.Sprintf("%d%d%d", year, month, day)
	fmt.Println("year-month-day:", year_month_day)
	fmt.Println("yearmonthday:", yearmonthday)

	y2019 := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	y2010 := time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)
	isOverTime := compareTime(y2019, y2010)
	fmt.Println("is after: ", isOverTime)
	get_now_time := GetNowTime()
	fmt.Println("get_now_time", get_now_time)
	time.Now().Date()
}

// 判断时间是否在某个时间之后
func compareTime(t1, t2 time.Time) bool {

	return t1.After(t2)
}

func GetNowTime() string {

	year, month, day := time.Now().Date()
	year_month_day := fmt.Sprintf("%d-%d-%d", year, month, day)
	return year_month_day
}

func GetNowDay() int {

	return time.Now().Day()
}

// 标准库把月份定义了一个枚举，既是英文也是int
func GetNowMonth() time.Month {

	return time.Now().Month()
}

func GetNowYear() int {

	return time.Now().Year()
}
