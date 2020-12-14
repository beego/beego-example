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
	"encoding/json"
	"strings"

	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
)

func main() {

	ctrl := &MainController{}

	web.BConfig.CopyRequestBody = true

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
	Name     string                 `json:"name" valid:"Required;Match(/^Bee.*/)"`
	Password string                 `json:"password"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *user) Valid(v *validation.Validation) {
	if strings.Index(u.Name, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Name", "名称里不能含有 admin")
	}
}

// curl --location --request POST 'localhost:8080/' \
// --header 'Content-Type: application/json' \
// --data-raw '{"name":"Beeadmin","password":"1234","metadata":{"phone":"12423434"},"created_at":"test"}'

// address: http://localhost:8080 Post
func (ctrl *MainController) Post() {
	valid := validation.Validation{}
	input := user{}

	if err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &input); err != nil {
		ctrl.Data["json"] = err.Error()
	}
	b, err := valid.Valid(&input)
	if err != nil {
		// handle error
	}
	errs := make(map[string]string)
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			errs[err.Key] = err.Message
		}
	}
	ctrl.Data["json"] = errs
	ctrl.ServeJSON()
}
