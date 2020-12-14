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
	"net/http"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func main() {

	web.Router("/", &MainController{})
	// The middleware definition is func(http.Handler) http.Handler
	web.RunWithMiddleWares(":8080", customMiddleware1, func(next http.Handler) http.Handler {
		return &customMiddleware2{
			next: next,
		}
	})

	// start the server and then request GET http://localhost:8080/
	// and then you will see:
	// log: in the custom middleware 1 from: /
	// log: in the custom middleware 2 from: /
}

// customMiddleware1 is a simple way to custom your middleware
// don't forget invoke next handler
func customMiddleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logs.Info("log: in the custom middleware 1 from: %s\n", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// customMiddleware2: usually you should think about using singleton design pattern
// and don't forget invoke next handler
type customMiddleware2 struct {
	next http.Handler
}

func (c *customMiddleware2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logs.Info("log: in the custom middleware 2 from: %s\n", r.RequestURI)
	c.next.ServeHTTP(w, r)
}

type MainController struct {
	web.Controller
}

func (m *MainController) Get() {
	m.Ctx.WriteString(fmt.Sprintf("hello world"))
}
