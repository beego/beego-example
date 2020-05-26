package main

import (
	"fmt"

	"github.com/astaxie/beego/httplib"
)

func main() {
	resp, err := httplib.Get("http://beego.me/").Debug(true).Response()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}
