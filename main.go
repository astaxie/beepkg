package main

import (
	"beepkg/controllers"
	"beepkg/controllers/admin"
	"github.com/astaxie/beego"
)

func main() {
	beego.BeeApp.RegisterController("/", &controllers.MainController{})
	beego.BeeApp.RegisterController("/admin", &admin.UserController{})
	beego.BeeApp.RegisterController("/admin/index", &admin.ArticleController{})
	beego.BeeApp.RegisterController("/admin/addpkg", &admin.AddController{})
	beego.BeeApp.RegisterController("/admin/editpkg/:id([0-9]+)", &admin.EditController{})
	beego.BeeApp.RegisterController("/admin/delpkg/:id([0-9]+)", &admin.DelController{})
	beego.BeeApp.RegisterController("/:pkg(.*)", &controllers.MainController{})
	beego.BeeApp.Run()
}
