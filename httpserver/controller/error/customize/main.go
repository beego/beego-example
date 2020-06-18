package main

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Ctx.WriteString("customize page not found")
	//c.Data= "page not found"
	//c.TplName = "404.tpl"
}

func (c *ErrorController) Error501() {
	c.Ctx.WriteString("customize server error")
	//c.Data["content"] = "server error"
	//c.TplName = "501.tpl"
}

func main() {
	beego.ErrorController(&ErrorController{}) // register error
	beego.Run()
}
