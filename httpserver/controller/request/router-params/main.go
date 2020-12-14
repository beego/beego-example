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
	web.Router("/:id", ctrl)

	web.Run()
}

// MainController:
// The controller must implement ControllerInterface
// Usually we extends web.Controller
type MainController struct {
	web.Controller
}

//curl --location --request GET 'localhost:8080/123'

// address: http://localhost:8080 GET
func (ctrl *MainController) Get() {

	id := ctrl.Ctx.Input.Param(":id")

	ctrl.Ctx.WriteString("Your router param ID is:" + id)
}
