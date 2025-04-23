package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Custom() {
	c.Data["Website"] = "localhost"
	c.Data["Email"] = "student@kbtu.kz"
	c.TplName = "index.tpl"
}
