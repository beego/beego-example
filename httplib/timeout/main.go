package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/httplib"
)

func main() {
	resp, err := httplib.Get("http://beego.me/").SetTimeout(100*time.Second, 30*time.Second).Response()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}
