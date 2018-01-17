package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type IndexController struct {
	//ControllerInterface
	beego.Controller
}

func init()  {
	fmt.Println("indexController init ... " )
}

func (this *IndexController) Prepare() {
	fmt.Println("this indexController prepare ... ")
}

func (this * IndexController) Get()  {
	fmt.Println("get func " )
 	this.Ctx.Redirect(404,"request error")
}