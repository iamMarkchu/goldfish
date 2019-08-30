package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/iamMarkchu/goldfish/tools"
)

var err error

func init()  {
	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	err = orm.RegisterDataBase("default", "mysql", "root:root@/gf_base?charset=utf8")
	go CheckError(err, "[RegisterDataBase Error]")
	orm.RegisterModel(new(User), new(Category))
	err = orm.RunSyncdb("default", false, true)
	go CheckError(err, "[RunSyncdb Error]")
	orm.Debug, err = beego.AppConfig.Bool("ormdebug")
	go CheckError(err, "[orm Debug Error]")
}
