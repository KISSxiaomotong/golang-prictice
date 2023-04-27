package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间戳
	var nowTime = time.Now()
	var timestamp = nowTime.Unix()
	fmt.Printf("当前时间戳为: %d\n", timestamp)

	//格式化输出时间
	var t = nowTime.UTC()
	var date = t.Format("2006-01-02 15:04:05")
	fmt.Printf("当前时间为: %s\n", date)
}
