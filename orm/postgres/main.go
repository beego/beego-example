package main

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

// Users -
type Users struct {
	ID   int    `orm:"column(id)"`
	Name string `orm:"column(name)"`
}

func init() {
	// set default database
	orm.RegisterDriver("postgres", orm.DRPostgres)

	// set default database
	orm.RegisterDataBase("default", "postgres", "postgres://username:passowrd@localhost/inesaim_test?sslmode=disable")

	// register model
	orm.RegisterModel(new(Users))

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	user := Users{Name: "Tom"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Name = "Jerry"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := Users{ID: user.ID}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
