package controllers

import (
	"blog/models"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	c.Data["IsAdmin"] = false
	user := c.GetSession("user")
	if user == nil {
		c.Data["IsLogin"] = false
	} else {
		data := user.(models.Users)
		c.Data["IsLogin"] = true
		c.Data["UserName"] = data.Name
		c.Data["UserAvatar"] = data.Avatar
		if data.IsAdmin == 1 {
			c.Data["IsAdmin"] = true
		}
	}

}
