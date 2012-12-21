package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beepkg/models"
)

type AddController struct {
	beego.Controller
}

func (this *AddController) Prepare() {

}

func (this *AddController) Get() {
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/add.tpl"
}

func (this *AddController) Post() {
	//数据处理
	this.Ctx.Request.ParseForm()
	pkgname := this.Ctx.Request.Form.Get("pkgname")
	content := this.Ctx.Request.Form.Get("content")
	beego.Info(this.Ctx.Request.Form)
	pk := models.GetCruPkg(pkgname)
	if pk.Id == 0 {
		var pp models.PkgEntity
		pp.Pid = 0
		pp.Pathname = pkgname
		pp.Intro = pkgname
		models.InsertPkg(pp)
		pk = models.GetCruPkg(pkgname)
	}
	var at models.Article
	at.Pkgid = pk.Id
	at.Content = content
	models.InsertArticle(at)
	this.Ctx.Redirect(302, "/admin/index")
}
