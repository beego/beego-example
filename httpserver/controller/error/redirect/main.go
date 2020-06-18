package main

import (
	"net/http"

	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &MainController{})
	beego.Router("/redirect", &RedirectController{})
	beego.Run()
}

type RedirectController struct {
	beego.Controller
}

func (r *RedirectController) Post() {
	// redirect to /, must use get method
	r.Redirect("/", http.StatusSeeOther)
}

type MainController struct {
	beego.Controller
}

func (m *MainController) Get() {
	m.Ctx.WriteString("hello world")
}
