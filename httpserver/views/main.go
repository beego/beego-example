// Copyright 2020 web-dev
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
	"github.com/beego/beego/v2/server/web"
)

func main() {

	ctrl := &MainController{}

	// GET http://localhost:8080/health => ctrl.Health()
	web.Router("/hello", ctrl, "get:Hello")

	web.Run()
}

// MainController:
// The controller must implement ControllerInterface
// Usually we extends web.Controller
type MainController struct {
	web.Controller
}

// address: http://localhost:8080/hello GET
func (ctrl *MainController) Hello() {

	// web-example/views/hello_world.html
	ctrl.TplName = "hello_world.html"
	ctrl.Data["name"] = "Hello()"

	// don't forget this
	_ = ctrl.Render()
}
