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
	"time"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
)

func main() {
	// create cache
	bm, err := cache.NewCache("redis", `{"key":"default", "conn":":6379", "password":"123456", "dbNum":"0"}`)
	if err != nil {
		logs.Error(err)
	}

	// put
	isPut := bm.Put("astaxie", 1, time.Second*10)
	logs.Info(isPut)

	isPut = bm.Put("hello", "world", time.Second*10)
	logs.Info(isPut)

	// get
	result := bm.Get("astaxie")
	logs.Info(string(result.([]byte)))

	multiResult := bm.GetMulti([]string{"astaxie", "hello"})
	for i := range multiResult {
		logs.Info(string(multiResult[i].([]byte)))
	}

	// isExist
	isExist := bm.IsExist("astaxie")
	logs.Info(isExist)

	// delete
	isDelete := bm.Delete("astaxie")
	logs.Info(isDelete)
}
