package main

import (
	"github.com/astaxie/beego"
	_ "github.com/iamMarkchu/goldfish/models"
	_ "github.com/iamMarkchu/goldfish/routers"
	_ "github.com/iamMarkchu/goldfish/tools/permission"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
