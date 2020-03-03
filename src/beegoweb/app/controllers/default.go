package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// @router /block [post]
func (ma *MainController) DealConsumeRecord() {
	ma.Ctx.WriteString("please")
}
