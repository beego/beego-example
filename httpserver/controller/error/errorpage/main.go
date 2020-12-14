package main

import (
	"html/template"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

func serverError(rw http.ResponseWriter, r *http.Request) {
	temp, _ := template.New("500.html").ParseFiles("/500.html")
	data := make(map[string]interface{})
	data["content"] = "Server Error"
	temp.Execute(rw, data)
}

func main() {
	web.ErrorHandler("500", serverError)
	web.Router("/", &MainController{})
	web.Run()
}

type MainController struct {
	web.Controller
}

func (m *MainController) Get() {
	m.Abort("serverError")
}
