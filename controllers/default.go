package controllers

import (
	"beepkg/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

var welcome = `
# Go pkg中文入门指南

`

func (this *MainController) Get() {
	crupkg := models.GetCruPkg(this.Ctx.Params[":pkg"])
	this.Data["CruPkg"] = crupkg
	pkglist := models.GetAllPkg()
	pm := make([]map[string]interface{}, len(pkglist))
	for _, pk := range pkglist {
		m := make(map[string]interface{}, 2)
		m["PKG"] = pk
		if this.Ctx.Params[":pkg"] == pk.Pathname {
			m["Cru"] = true
		} else {
			m["Cru"] = false
		}
		pm = append(pm, m)
	}
	this.Data["PkgList"] = pm
	if crupkg.Id == 0 {
		this.Data["Content"] = welcome //template.HTML(string(blackfriday.MarkdownCommon([]byte(welcome))))
	} else {
		at := models.GetArticle(crupkg.Id)
		this.Data["Content"] = at.Content //template.HTML(string(blackfriday.MarkdownCommon([]byte(at.Content))))
	}
	this.TplNames = "index.tpl"
}
