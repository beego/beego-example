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

	// we register the path / to &MainController
	// if we don't pass methodName as third param
	// web will use the default mappingMethods
	// GET http://localhost:8080  -> Get()
	// POST http://localhost:8080 -> Post()
	// ...
	web.Router("/", ctrl)

	// GET http://localhost:8080/health => ctrl.Health()
	web.Router("/health", ctrl, "get:Health")

	// POST http://localhost:8080/update => ctrl.Update()
	web.Router("/update", ctrl, "post:Update")

	// support multiple http methods.
	// POST or GET http://localhost:8080/update => ctrl.GetOrPost()
	web.Router("/getOrPost", ctrl, "get,post:GetOrPost")

	// support any http method
	// POST, GET, PUT, DELETE... http://localhost:8080/update => ctrl.Any()
	web.Router("/any", ctrl, "*:Any")

	web.Run()
}

// MainController:
// The controller must implement ControllerInterface
// Usually we extends web.Controller
type MainController struct {
	web.Controller
}

// address: http://localhost:8080 GET
func (ctrl *MainController) Get() {

	// web-example/views/hello_world.html
	ctrl.TplName = "hello_world.html"
	ctrl.Data["name"] = "Get()"

	// don't forget this
	_ = ctrl.Render()
}

// GET http://localhost:8080/health
func (ctrl *MainController) Health() {
	// web-example/views/hello_world.html
	ctrl.TplName = "hello_world.html"
	ctrl.Data["name"] = "Health()"

	// don't forget this
	_ = ctrl.Render()
}

// POST http://localhost:8080/update
func (ctrl *MainController) Update() {
	// web-example/views/hello_world.html
	ctrl.TplName = "hello_world.html"

	ctrl.Data["name"] = "Update()"

	// don't forget this
	_ = ctrl.Render()
}

// GET or POST http://localhost:8080/update
func (ctrl *MainController) GetOrPost() {
	// web-example/views/hello_world.html
	ctrl.TplName = "hello_world.html"

	ctrl.Data["name"] = "GetOrPost()"

	// don't forget this
	_ = ctrl.Render()
}

// any http method http://localhost:8080/any
func (ctrl *MainController) Any() {
	// web-example/views/hello_world.html
	ctrl.TplName = "hello_world.html"

	ctrl.Data["name"] = "Any()"

	// don't forget this
	_ = ctrl.Render()
}
