package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//根据当前时间打印月份
	var nowTime string
	var nowMonth string
	nowTime = time.Now().UTC().Format("2006-01-02 15:04:05")

	month, _ := strconv.Atoi(string([]rune(nowTime)[5:7]))

	switch month {
	case 1:
		nowMonth = "一月份"
	case 2:
		nowMonth = "二月份"
	case 3:
		nowMonth = "三月份"
	case 4:
		nowMonth = "四月份"
	case 5:
		nowMonth = "五月份"
	case 6:
		nowMonth = "六月份"
	case 7:
		nowMonth = "七月份"
	case 8:
		nowMonth = "八月份"
	case 9:
		nowMonth = "九月份"
	case 10:
		nowMonth = "十月份"
	case 11:
		nowMonth = "十一月份"
	case 12:
		nowMonth = "十二月份"
	}
	fmt.Printf(nowMonth)
}
