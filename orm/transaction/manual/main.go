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
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `orm:"column(id)"`
	Name string `orm:"column(name)"`
}

func init() {
	// need to register models in init
	orm.RegisterModel(new(User))

	// need to register db driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// need to register default database
	orm.RegisterDataBase("default", "mysql", "beego:test@tcp(192.168.0.105:13306)/orm_test?charset=utf8")

	orm.RunSyncdb("default", true, false)
}

func main() {
	o := orm.NewOrm()
	to, err := o.Begin()
	if err != nil {
		logs.Error("start the transaction failed")
		return
	}

	user := new(User)
	user.Name = "test_transaction"

	// do something with to. to is an instance of TxOrm

	// insert data
	// Using txOrm to execute SQL
	_, err = to.Insert(user)

	if err != nil {
		logs.Error("execute transaction's sql fail, rollback.", err)
		err = to.Rollback()
		if err != nil {
			logs.Error("roll back transaction failed", err)
		}
	} else {
		err = to.Commit()
		if err != nil {
			logs.Error("commit transaction failed.", err)
		}
	}

}
