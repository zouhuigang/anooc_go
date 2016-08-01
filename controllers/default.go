package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "anooc.com"
	c.Data["Email"] = "952750120@qq.com"
	c.TplName = "index.tpl"
}




