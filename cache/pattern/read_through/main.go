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
	"fmt"
	"github.com/beego/beego/v2/client/cache"
	"time"
)

func main() {
	var c cache.Cache = cache.NewMemoryCache()
	var err error
	c, err = cache.NewReadThroughCache(c,
		// expiration, same as the expiration of key
		time.Minute,
		// load func, how to load data if the key is absent.
		// in general, you should load data from database.
		func(ctx context.Context, key string) (any, error) {
			return fmt.Sprintf("hello, %s", key), nil
		})
	if err != nil {
		panic(err)
	}

	val, err := c.Get(context.Background(), "Beego")
	if err != nil {
		panic(err)
	}
	// print hello, Beego
	fmt.Print(val)
}
