package main

import (
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/abort", &AbortController{})
	beego.Run()
}

type AbortController struct {
	beego.Controller
}

func (a *AbortController) Get() {
	a.Abort("401")
	// None of the following
	user := a.GetString("user")
	if user != "" {
		a.Ctx.WriteString("hello world")
		return
	}
}
