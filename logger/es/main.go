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
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/beego/beego/v2/core/logs/es"
)

func main() {

	// don't forget import es package
	// this statement:
	// _ "github.com/beego/beego/v2/adapter/logs/es"

	// before you run this example, please start ES server and update address to your real address
	err := logs.SetLogger(logs.AdapterEs, `{"dsn":"http://localhost:9200/","level":1}`)
	if err != nil {
		panic(err)
	}
	logs.Info("hello beego")
}
