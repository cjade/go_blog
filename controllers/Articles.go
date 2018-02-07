package controllers

type ArticlesController struct {
	BaseController
}

func (c *ArticlesController) Details() {
	c.Ctx.Input.Param(":id")
	c.Data["PageTitle"] = "文章详情"
	c.Layout = "home/public/layout.html"
	c.TplName = "home/articleDetails.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Banner"] = "home/public/_banner.html"
}
