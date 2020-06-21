package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.Info("auto load config name is", beego.AppConfig.String("name"))
}
