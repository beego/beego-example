package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// User -
type User struct {
	ID   int    `orm:"column(id)"`
	Name string `orm:"column(name)"`
}

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/beego?charset=utf8")
}

func main() {
	orm.RunSyncdb("default", false, true)

	o := orm.NewOrm()
	o.Using("default")

	user := new(User)
	user.Name = "mike"

	fmt.Println(o.Insert(user))
}
