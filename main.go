package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beepkg/controllers"
	"github.com/astaxie/beepkg/controllers/admin"
)

func main() {
	beego.RegisterController("/", &controllers.MainController{})
	beego.RegisterController("/admin", &admin.UserController{})
	beego.RegisterController("/admin/index", &admin.ArticleController{})
	beego.RegisterController("/admin/addpkg", &admin.AddController{})
	beego.RegisterController("/admin/editpkg/:id([0-9]+)", &admin.EditController{})
	beego.RegisterController("/admin/delpkg/:id([0-9]+)", &admin.DelController{})
	beego.RegisterController("/:pkg(.*)", &controllers.MainController{})
	beego.Run()
}
