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
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/client/httplib/filter/opentracing"
	"github.com/beego/beego/v2/core/logs"
)

func main() {

	// don't forget this to inject the opentracing API's implementation
	// opentracing2.SetGlobalTracer()

	builder := opentracing.FilterChainBuilder{}
	req := httplib.Get("http://beego.me/")
	// only work for this request, or using SetDefaultSetting to support all requests
	req.AddFilters(builder.FilterChain)

	resp, err := req.Response()
	if err != nil {
		logs.Error("could not get response: ", err)
	} else {
		logs.Info(resp)
	}
}
