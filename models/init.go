package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func Init() {
	db_username := beego.AppConfig.String("DB_USERNAME")
	db_password := beego.AppConfig.String("DB_PASSWORD")
	db_host := beego.AppConfig.String("DB_HOST")
	db_port := beego.AppConfig.String("DB_PORT")
	db_database := beego.AppConfig.String("DB_DATABASE")
	db_prefix := beego.AppConfig.String("DB_PREFIX")
	if db_port == "" {
		db_port = "3306"
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", db_username+":"+db_password+"@tcp("+db_host+":"+db_port+")/"+db_database+"?charset=utf8", 30)
	orm.RegisterModelWithPrefix(db_prefix,
		new(Articles),
		new(Labels),
		new(Users))
	// create table
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}

func TableName(name string) string {
	return beego.AppConfig.String("DB_PREFIX") + name
}
