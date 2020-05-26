package main

import (
	"io/ioutil"
	"log"

	"github.com/astaxie/beego/httplib"
)

func main() {
	// Upload File
	fileReq := httplib.Post("http://beego.me/")
	fileReq.Param("username", "astaxie")
	fileReq.Param("password", "123456")
	fileReq.PostFile("uploadfile", "hello.txt")

	// Bigfile
	bigFileReq := httplib.Post("http://beego.me/")
	bt, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		log.Fatal("read file err:", err)
	}
	bigFileReq.Body(bt)
}
