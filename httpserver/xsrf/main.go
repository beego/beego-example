// Copyright 2020 
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"html/template"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.BConfig.WebConfig.EnableXSRF = true
	web.BConfig.WebConfig.XSRFKey = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	web.BConfig.WebConfig.XSRFExpire = 3600

	mc := &MainController{}

	web.Router("/xsrfpage", mc, "get:XsrfPage")
	web.Router("/new_message", mc, "post:NewMessage")

	web.Run(":8080")
}

type MainController struct {
	web.Controller
}

func (mc *MainController) XsrfPage() {
	mc.XSRFExpire = 7200
	mc.Data["xsrfdata"] = template.HTML(mc.XSRFFormHTML())
	mc.TplName = "xsrf.html"
}

func (mc *MainController) NewMessage() {
	v, _ := mc.Input()
	mc.Ctx.WriteString("hello" + v.Get("message"))
}
