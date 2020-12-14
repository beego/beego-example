// Copyright 2020 astaxie
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

	"github.com/beego/beego/v2/server/web"
	apmbeego "github.com/opentracing-contrib/beego"
)

func main() {
	web.Router("/", &MainController{})
	// Actually, I just use the opentracing-contrib/web as example but I do not check
	// whether it is a good middleware
	web.RunWithMiddleWares("localhost:8080", apmbeego.Middleware("bee-go-demo"))

	// start the server and then request GET http://localhost:8080/
}

type MainController struct {
	web.Controller
}

func (m *MainController) Get() {
	m.Ctx.WriteString(fmt.Sprintf("hello world"))
}
