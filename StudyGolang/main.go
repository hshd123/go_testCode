package main

import (
	_ "StudyGolang/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
 )

func main() {
	beego.BeeLogger.EnableFuncCallDepth(true)
	s1 := `{"filename":"studygolang.log","level":7,"maxlines":1000024,"maxsize":10000024,"daily":true,"maxdays":1000024}`
    beego.BeeLogger.EnableFuncCallDepth(true)
    beego.BeeLogger.SetLogFuncCallDepth(7)
    beego.BeeLogger.SetLogger(logs.AdapterFile,s1)
 	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
 	}
	beego.Run()
}
