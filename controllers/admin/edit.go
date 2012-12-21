package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beepkg/models"
	"strconv"
)

type EditController struct {
	beego.Controller
}

func (this *EditController) Prepare() {

}

func (this *EditController) Get() {
	var at models.Article
	atId, _ := strconv.Atoi(this.Ctx.Params[":id"])
	at = models.GetArticleById(atId)
	pk := models.GetPkgByID(at.Pkgid)
	this.Data["article"] = at
	this.Data["pkg"] = pk
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/edit.tpl"
}

func (this *EditController) Post() {
	//数据处理
	this.Ctx.Request.ParseForm()
	id := this.Ctx.Request.Form.Get("id")
	atid, _ := strconv.Atoi(id)
	at := models.GetArticleById(atid)
	content := this.Ctx.Request.Form.Get("content")
	at.Content = content
	models.InsertArticle(at)
	this.Ctx.Redirect(302, "/admin/index")
}
