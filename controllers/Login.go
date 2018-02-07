package controllers

import (
	"fmt"

	"blog/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Login() {
	c.Data["PageTitle"] = "登录"
	c.Layout = "home/public/layout.html"
	c.TplName = "home/login.html"
}

func (c *LoginController) DoLogin() {
	email := c.Input().Get("email")
	password := c.Input().Get("password")
	c.Ctx.WriteString(fmt.Sprint(email, password))

	return
}

func (c *LoginController) Register() {
	//flash 数据
	flash := beego.ReadFromRequest(&c.Controller)

	c.Data["PageTitle"] = "注册"
	c.Layout = "home/public/layout.html"
	c.TplName = "home/register.html"
	if _, ok := flash.Data["error"]; ok {
		// 显示设置成功
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Noty"] = "home/public/_toastr_error.html"
	}
}

func (c *LoginController) DoRegister() {
	name := c.Input().Get("name")
	email := c.Input().Get("email")
	password := c.Input().Get("password")
	password_confirmation := c.Input().Get("password_confirmation")
	flash := beego.NewFlash()
	if name == "" {
		flash.Error("用户名不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
		return
	}
	if email == "" {
		flash.Error("邮箱不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)

		return
	}
	if password != password_confirmation {
		flash.Error("两次密码不一致")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)

		return
	}

	users := models.GetUserByEmail(email)
	if users.Email != "" {
		c.Ctx.WriteString(fmt.Sprint("邮箱已存在"))
		return
	}

	var data models.Users
	data.Name = name
	data.Email = email
	data.Password = password
	_, err := models.CreateUser(data)
	if err == nil {
		c.SetSession("user", models.GetUserByEmail(email))
		c.Redirect("/", 301)
	} else {
		c.Data["json"] = map[string]interface{}{"code": -1, "message": "添加出错"}
		c.ServeJSON()
	}
}
