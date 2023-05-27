// Copyright 2023 beego-dev
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
	c := cache.NewMemoryCache()
	c, err := cache.NewBloomFilterCache(c, func(ctx context.Context, key string) (any, error) {
		return fmt.Sprintf("hello, %s", key), nil
	}, &AlwaysExist{}, time.Minute)
	if err != nil {
		panic(err)
	}

	val, err := c.Get(context.Background(), "Beego")
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}

type AlwaysExist struct {
}

func (a *AlwaysExist) Test(data string) bool {
	return true
}

func (a *AlwaysExist) Add(data string) {

}
