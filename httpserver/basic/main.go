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
	"github.com/beego/beego/v2/server/web"
)

func main() {
	// now you start the beego as http server.
	// it will listen to port 8080
	web.Run()

	// it will listen to 8080
	// beego.Run("localhost")

	// it will listen to 8089
	// beego.Run(":8089")

	// it will listen to 8089
	// beego.Run("127.0.0.1:8089")
}
