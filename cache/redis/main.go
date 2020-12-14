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
	"context"
	"time"

	"github.com/beego/beego/v2/client/cache"
	// don't forget this
	_ "github.com/beego/beego/v2/client/cache/redis"
	"github.com/beego/beego/v2/core/logs"
)

func main() {
	// create cache
	bm, err := cache.NewCache("redis", `{"key":"default", "conn":"192.168.0.105:6379", "dbNum":"0"}`)
	if err != nil {
		logs.Error(err)
	}

	// put
	isPut := bm.Put(context.Background(), "astaxie", 1, time.Second*10)
	logs.Info(isPut)

	isPut = bm.Put(context.Background(), "hello", "world", time.Second*10)
	logs.Info(isPut)

	// get
	result, _ := bm.Get(context.Background(), "astaxie")
	logs.Info(string(result.([]byte)))

	multiResult, _ := bm.GetMulti(context.Background(), []string{"astaxie", "hello"})
	for i := range multiResult {
		logs.Info(string(multiResult[i].([]byte)))
	}

	// isExist
	isExist, _ := bm.IsExist(context.Background(), "astaxie")
	logs.Info(isExist)

	// delete
	isDelete := bm.Delete(context.Background(), "astaxie")
	logs.Info(isDelete)
}
