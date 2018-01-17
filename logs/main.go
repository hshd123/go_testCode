package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

func main() {
	l1 := logs.NewLogger(10000)
	l1.SetLevel(logs.LevelDebug)
	l1.EnableFuncCallDepth(true) //显示行号
	//logs.Async(1024 * 1024 * 1024)
	//for i := 0; i < 10; i++ {
	//	l1.Info("i = ", i)
	//	logs.Error("i = ", i)
	//}
	//l1.Debug("this is debug message")
	l1.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/test1.log","level":7,"maxsize":10000000000,"daily":true,"maxdays":10000000}`)
	l1.Debug("sss")
	for i := 0; i < 1000000; i++ {
		str := "log ---- " + string(time.Now().String())
		l1.Debug(str)
		fmt.Println(str)
		time.Sleep(time.Second * 3)
	}

}
