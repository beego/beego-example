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
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/captcha"
)

// init captcha
var cpt *captcha.Captcha

// Controller -
type Controller struct {
	beego.Controller
}

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 100
	cpt.StdHeight = 40
}

func main() {
	ctrl := &Controller{}

	beego.Router("/", ctrl)
	beego.Router("/sendCaptcha", ctrl, "post:Captcha")

	beego.Run()
}

// Get - address: http://127.0.0.1:8080/
func (ctrl *Controller) Get() {
	ctrl.TplName = "captcha.html"
	ctrl.Data["name"] = "Home"

	// don't forget this
	_ = ctrl.Render()
}

// Captcha - address: http://127.0.0.1:8080/sendCaptcha
func (ctrl *Controller) Captcha() {
	ctrl.TplName = "captcha.html"

	if !cpt.VerifyReq(ctrl.Ctx.Request) {
		logs.Error("Captcha does not match")
		_ = ctrl.Render()
		return
	}

	logs.Info("matched")
	_ = ctrl.Render()
}
