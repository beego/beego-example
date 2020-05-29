// Copyright 2020 beego-dev
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
	"github.com/astaxie/beego"
)

// to test this
// you should browse POST http://localhost:8080/session to set session firstly
func main() {

	// create contr
	ctrl := &MainController{}

	// POST http://localhost:8080/cookie => ctrl.PutCookie()
	beego.Router("/cookie", ctrl, "post:PutCookie")

	// GET http://localhost:8080/cookie => ctrl.ReadCookie()
	beego.Router("/cookie", ctrl, "get:ReadCookie")

	beego.Run()
}

type MainController struct {
	beego.Controller
}

func (ctrl *MainController) PutCookie()  {
	// put something into cookie
	ctrl.Ctx.SetCookie("name", "beego cookie")

	// beego-example/views/hello_world.html
	ctrl.TplName = "hello_world.html"
	ctrl.Data["name"] = "PutCookie"
	_ = ctrl.Render()
}

func (ctrl *MainController) ReadCookie()  {
	// beego-example/views/hello_world.html
	ctrl.TplName = "hello_world.html"
	ctrl.Data["name"] = ctrl.Ctx.GetCookie("name")
	// don't forget this
	_ = ctrl.Render()
}

