package main

import (
	_ "blog/routers"

	"blog/models"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.Init()
}

func main() {
	beego.Run()
}
