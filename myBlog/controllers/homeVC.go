package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type HomeController struct {
	beego.Controller
}

func (this * HomeController) Get()  {
	fmt.Println("get http")
	this.TplName = "home.html"
}