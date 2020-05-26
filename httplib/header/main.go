package main

import (
	"fmt"

	"github.com/astaxie/beego/httplib"
)

func main() {
	req := httplib.Post("http://beego.me/")
	req.Header("Accept-Encoding", "gzip,deflate,sdch")
	req.Header("Host", "beego.me")

	header := req.GetRequest().Header
	fmt.Println(header)
}
