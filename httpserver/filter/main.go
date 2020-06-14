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
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {

	ctrl := &MainController{}
	//定义路由过滤器
	beego.InsertFilter("/user/*", beego.BeforeExec, filterFunc)
	// we register the path / to &MainController
	// if we don't pass methodName as third param
	// beego will use the default mappingMethods
	// GET http://localhost:8080  -> Get()
	// POST http://localhost:8080 -> Post()
	// ...
	beego.Router("/", ctrl)
	//用户登录
	beego.Router("/login", ctrl, "get:Login")
	//用户中心个人信息页面展示
	beego.Router("/user/info", ctrl, "get:User")

	beego.Run()
}

// MainController:
// The controller must implement ControllerInterface
// Usually we extends beego.Controller
type MainController struct {
	beego.Controller
}

// address: http://localhost:8080 GET
func (ctrl *MainController) Get() {

	// beego-example/views/hello_world.html
	ctrl.TplName = "hello_world.html"
	ctrl.Data["name"] = "Get()"

	// don't forget this
	_ = ctrl.Render()
}

// any http method http://localhost:8080/login
func (ctrl *MainController) Login() {
	fmt.Println("过滤校验login")
	// beego-example/views/hello_world.html
	ctrl.TplName = "hello_world.html"

	ctrl.Data["name"] = "Login()"

	// don't forget this
	_ = ctrl.Render()
}

// any http method http://localhost:8080/user/info
func (ctrl *MainController) User() {
	// beego-example/views/hello_world.html
	ctrl.TplName = "hello_world.html"

	ctrl.Data["name"] = "hello beego"

	// don't forget this
	_ = ctrl.Render()
}

/* 定义过滤函数 */
func filterFunc(ctx *context.Context) {
	//过滤校验
	fmt.Println("过滤校验")
	var id int
	_ = ctx.Input.Bind(&id, "id") //id ==123
	fmt.Println(id)
	if id == 0 {
		ctx.Redirect(302, "/login")
		return
	}
}
