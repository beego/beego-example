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

	web.Run()
}

// MainController:
// The controller must implement ControllerInterface
// Usually we extends web.Controller
type MainController struct {
	web.Controller
}

type user struct {
	Name     string `form:"name"`
	Password string `form:"password" json:"custom_password"`
}

// curl --location --request POST 'localhost:8080/' \
// --form 'name=test' \
// --form 'password=1234'

// address: http://localhost:8080 Post
func (ctrl *MainController) Post() {
	input := user{}

	if err := ctrl.ParseForm(&input); err != nil {
		ctrl.Ctx.WriteString(err.Error())
		return
	}
	ctrl.Data["json"] = input
	ctrl.ServeJSON()
}
