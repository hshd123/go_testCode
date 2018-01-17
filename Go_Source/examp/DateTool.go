package main

import (
	"fmt"
	"time"
)

//将时间戳转换成字符串
// layout 2006-01-02 15:04:05 基准时间
func SecToTimeStr(timeStamp int64 , layout string) string  {
	if timeStamp == 0 {
		return ""
	}
	t1 := time.Unix(timeStamp , 0)
	retStr := t1.Format(layout)
	return retStr
}

//将字符串时间戳转换成秒
// layout 2006-01-02 15:04:05 基准时间
func TimeStrToSec(timeStr string , layout string) int64  {
	t1 ,err := time.Parse(layout , timeStr)
	if err != nil {
		return 0
	}
	return t1.Unix()
}

func Hello()  {
	fmt.Println("hello world")
}