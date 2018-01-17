package main

import (
	"github.com/astaxie/beego"
	_ "myBlog/routers"
	"myBlog/models"
)

func init()  {
	models.RegisterDB()
}

func main() {
	beego.Run()
}
