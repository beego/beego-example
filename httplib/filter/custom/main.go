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
	"context"
	"net/http"
	"time"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
)

// implement custom filter
func main() {
	httplib.SetDefaultSetting(httplib.BeegoHTTPSettings{

		FilterChains: []httplib.FilterChain{
			myFilter,
		},

		UserAgent:        "beegoServer",
		ConnectTimeout:   60 * time.Second,
		ReadWriteTimeout: 60 * time.Second,
		Gzip:             true,
		DumpBody:         true,
	})

	req := httplib.Get("http://beego.me")
	
	// or req.AddFilters(myFilter)


	resp, err := req.Response()
	if err != nil {
		logs.Error("could not get response: ", err)
	} else {
		logs.Info(resp)
	}
}

func myFilter(next httplib.Filter) httplib.Filter {
	return func(ctx context.Context, req *httplib.BeegoHTTPRequest) (*http.Response, error) {
		r := req.GetRequest()
		logs.Info("hello, here is the filter: ", r.URL)
		// Never forget invoke this. Or the request will not be sent
		return next(ctx, req)
	}
}
