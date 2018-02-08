package controllers

type IndexController struct {
	BaseController
}

/**
 *首页
 */
func (c *IndexController) Index() {
	c.Data["PageTitle"] = "首页"
	c.Layout = "home/public/layout.html"
	c.TplName = "home/index.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Banner"] = "home/public/_banner.html"
}
