package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin bool
}

func (c *BaseController) init() {
	user := c.GetSession("user")
	if user == nil {
		c.IsLogin = false
	} else {
		c.IsLogin = true
	}

}
