package main

import (
	"blog/models"
	_ "blog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("DB_USERNAME")+":"+beego.AppConfig.String("DB_PASSWORD")+"@tcp("+beego.AppConfig.String("DB_HOST")+":"+beego.AppConfig.String("DB_PORT")+")/"+beego.AppConfig.String("DB_DATABASE")+"?charset=utf8", 30)
	orm.RegisterModelWithPrefix("blog_", new(models.Articles), new(models.Labels), new(models.Users))
	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	orm.Debug = true
	beego.Run()
}
