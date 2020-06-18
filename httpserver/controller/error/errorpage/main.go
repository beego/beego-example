package main

import (
	"html/template"
	"net/http"

	"github.com/astaxie/beego"
)

func serverError(rw http.ResponseWriter, r *http.Request) {
	temp, _ := template.New("500.html").ParseFiles("/500.html")
	data := make(map[string]interface{})
	data["content"] = "Server Error"
	temp.Execute(rw, data)
}

func main() {
	beego.ErrorHandler("500", serverError)
	beego.Router("/", &MainController{})
	beego.Run()
}

type MainController struct {
	beego.Controller
}

func (m *MainController) Get() {
	m.Abort("serverError")
}
