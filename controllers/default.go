package controllers

import (
	"github.com/astaxie/beego"
)

// MainController : export
type MainController struct {
	beego.Controller
}

// Get : default get cnotroller
func (c *MainController) Get() {

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "welcome.html"

}
