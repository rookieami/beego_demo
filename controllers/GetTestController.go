package controllers

import (
	"github.com/astaxie/beego"
)

type GetTestController struct {
	beego.Controller
}

func (c *GetTestController) Get() {
	c.Ctx.WriteString("gettets")
}
