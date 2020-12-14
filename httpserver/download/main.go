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
	web.Router("/download", ctrl, "get:Download")
	web.Router("/download_file", ctrl, "get:DownloadFile")
	web.Run()
}

type MainController struct {
	web.Controller
}

func (ctrl *MainController) DownloadFile() {
	// The file LICENSE is under root path.
	// and the downloaded file name is license.txt
	ctrl.Ctx.Output.Download("LICENSE", "license.txt")
}

// Download is an example that download the content stored in memory
// when you access localhost:8080/download
// A file named "mytest.txt" with content "Hello" will be downloaded
func (ctrl *MainController) Download() {
	output := ctrl.Ctx.Output
	output.Header("Content-Disposition", "attachment;filename=mytest.txt;")
	output.Header("Content-Description", "File Transfer")
	output.Header("Content-Type", "application/octet-stream")
	output.Header("Content-Transfer-Encoding", "binary")
	output.Header("Expires", "0")
	output.Header("Cache-Control", "must-revalidate")
	output.Header("Pragma", "public")
	ctrl.Ctx.WriteString("Hello")
}
