package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "GET:Index")
	beego.Router("/article/:id([0-9]+)", &controllers.ArticlesController{}, "GET:Details")
	beego.Router("login", &controllers.LoginController{}, "GET:Login")
	beego.Router("doLogin", &controllers.LoginController{}, "POST:DoLogin")
	beego.Router("register", &controllers.LoginController{}, "GET:Register")
	beego.Router("doRegister", &controllers.LoginController{}, "POST:DoRegister")
}
