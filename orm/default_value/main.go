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
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/client/orm/filter/bean"

	// don't forget this
	_ "github.com/go-sql-driver/mysql"
)

// User -
type User struct {
	ID            int    `orm:"column(id)"`
	Name          string `orm:"column(name)"`
	Age           int    `default:"12"`
	AgeInOldStyle int    `orm:"default(13);bee()"`
}

func init() {
	// need to register models in init
	orm.RegisterModel(new(User))

	// need to register db driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// need to register default database
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/beego?charset=utf8")
}

func main() {
	builder := bean.NewDefaultValueFilterChainBuilder(nil, true, true)
	orm.AddGlobalFilterChain(builder.FilterChain)
	o := orm.NewOrm()
	_, _ = o.Insert(&User{
		ID: 1,
		Name: "Tom",
	})
}
