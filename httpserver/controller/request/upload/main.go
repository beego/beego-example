// Copyright 2020 
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
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	ctrl := &MainController{}

	// upload file by path /upload
	web.Router("/upload", ctrl, "post:Upload")
	web.Router("/upload", ctrl, "get:UploadPage")
	web.Router("/save", ctrl, "post:Save")

	web.Run()
}

// MainController:
// The controller must implement ControllerInterface
// Usually we extends web.Controller
type MainController struct {
	web.Controller
}

// GET http://localhost:8080/upload
// and you will see the upload page
func (ctrl *MainController) UploadPage() {
	ctrl.TplName = "upload.html"
}

// POST http://localhost:8080/upload
// you will see "success"
// and the file name (actual file name, not the key you use in GetFile
func (ctrl *MainController) Upload() {
	// key is the file name
	file, fileHeader, err := ctrl.GetFile("upload.txt")
	if err != nil {
		logs.Error("save file failed, ", err)
		ctrl.Ctx.Output.Body([]byte(err.Error()))
	} else {
		// don't forget to close
		defer file.Close()

		logs.Info(fileHeader.Filename)
		ctrl.Ctx.Output.Body([]byte("success"))
	}


}

// POST http://localhost:8080/save
// you will see the file /tmp/upload.txt and "success"
// and if you run this on Windows platform, don't forget to change the target file path
func (ctrl *MainController) Save() {
	err := ctrl.SaveToFile("save.txt", "/tmp/upload.txt")
	if err != nil {
		logs.Error("save file failed, ", err)
		ctrl.Ctx.Output.Body([]byte(err.Error()))
	} else {
		ctrl.Ctx.Output.Body([]byte("success"))
	}
}
