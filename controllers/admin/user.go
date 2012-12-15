package admin

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	this.Data["Username"] = "astaxie"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "admin/signin.tpl"
}
