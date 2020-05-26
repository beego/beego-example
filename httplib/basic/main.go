package main

import (
	"crypto/tls"
	"fmt"

	"github.com/astaxie/beego/httplib"
)

func main() {
	// Get
	req := httplib.Get("http://beego.me/")
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(str)

	// Post
	req = httplib.Post("http://beego.me/")
	req.Param("username", "astaxie")
	req.Param("password", "123456")

	// Https
	req = httplib.Get("http://beego.me/")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
}
