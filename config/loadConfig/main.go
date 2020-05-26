package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	ConfigFile = "./app.conf"
)

func main() {
	err := beego.LoadAppConfig("ini", ConfigFile)
	if err != nil {
		logs.Critical("An error occurred:", err)
		panic(err)
	}
	logs.Info("load config name is",beego.AppConfig.String("name"))
}
