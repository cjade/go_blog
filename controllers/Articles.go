package controllers

import (
	"github.com/astaxie/beego"
)

type ArticlesController struct {
	beego.Controller
}

func (c *ArticlesController) Details() {
	c.Ctx.Input.Param(":id")
	c.Data["PageTitle"] = "文章详情"
	c.Data["IsLogin"] = true
	c.Layout = "home/public/layout.html"
	c.TplName = "home/articleDetails.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Banner"] = "home/public/_banner.html"
}
