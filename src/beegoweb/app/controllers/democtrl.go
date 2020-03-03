package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// CMS API
type CMSController struct {
	beego.Controller
}

func (c *CMSController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
}

// @router /staticblock/:key [get]
func (this *CMSController) StaticBlock() {

}

// @Title AddWeapon
// @Description create a new weapon
// @Param	body		body 	models.Weapon	true "body for weapon content"
// @Success 200 {int} models.Weapon.Id
// @Failure 403 body is empty
// @router /all [get]
func (this *CMSController) AllBlock() {
	this.Ctx.Output.Body([]byte("AllBlock"))
	name := this.GetString("name")
	logs.Info(name)
	name2 := this.Input().Get("name")
	logs.Info(name2)
}
