package main

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
)

// User -
type User struct {
	ID   int    `orm:"column(id)"`
	Name string `orm:"column(name)"`
}

func init() {
	// need to register models in init
	orm.RegisterModel(new(User))

	// need to register db driver
	orm.RegisterDriver("sqlite3", orm.DRSqlite)

	// need to register default database
	orm.RegisterDataBase("default", "sqlite3", "beego.db")
}

func main() {
	// automatically build table
	orm.RunSyncdb("default", false, true)

	// create orm object
	o := orm.NewOrm()

	// data
	user := new(User)
	user.Name = "mike"

	// insert data
	o.Insert(user)
}
