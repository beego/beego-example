package main

import (
	"github.com/beego/beego/core/logs"
	"github.com/beego/beego/server/web"
)

func main() {
	// be careful about the work directory
	val, _ := web.AppConfig.String("name")
	logs.Info("auto load config name is", val)
}
