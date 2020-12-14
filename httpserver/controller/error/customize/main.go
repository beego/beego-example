package main

import (
	"github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	web.Controller
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
	web.ErrorController(&ErrorController{}) // register error
	web.Run()
}
