package main

import (
	"github.com/astaxie/beego/core/logs"
	"github.com/astaxie/beego/server/web"
)

var (
	ConfigFile = "./app.conf"
)

func main() {
	err := web.LoadAppConfig("ini", ConfigFile)
	if err != nil {
		logs.Critical("An error occurred:", err)
		panic(err)
	}

	val, _ := web.AppConfig.String("name")

	logs.Info("load config name is", val)
}
