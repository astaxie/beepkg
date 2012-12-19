package main

import (
	"beepkg/controllers"
	"beepkg/controllers/admin"
	"github.com/astaxie/beego"
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
