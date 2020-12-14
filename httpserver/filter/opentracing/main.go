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
	"time"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/opentracing"
)

func main() {
	// don't forget this to inject the opentracing API's implementation
	// opentracing2.SetGlobalTracer()

	web.BConfig.AppName = "my app"

	ctrl := &MainController{}
	web.Router("/hello", ctrl, "get:Hello")
	fb := &opentracing.FilterChainBuilder{}
	web.InsertFilterChain("/*", fb.FilterChain)
	web.Run(":8080")
	// after you start the server
}

type MainController struct {
	web.Controller
}

func (ctrl *MainController) Hello() {
	time.Sleep(time.Second)
	ctrl.Ctx.ResponseWriter.Write([]byte("Hello, world"))
}
