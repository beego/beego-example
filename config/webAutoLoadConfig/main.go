package main

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	// be careful about the work directory
	val, _ := web.AppConfig.String("name")
	logs.Info("auto load config name is", val)
}
