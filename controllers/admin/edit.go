package admin

import (
	"beepkg/models"
	"github.com/astaxie/beego"
	"strconv"
)

type EditController struct {
	beego.Controller
}

func (this *EditController) Prepare() {

}

func (this *EditController) Get() {
	var at models.Article
	atId, _ := strconv.Atoi(this.Ct.Params[":id"])
	at = models.GetArticleById(atId)
	pk := models.GetPkgByID(at.Pkgid)
	this.Data["article"] = at
	this.Data["pkg"] = pk
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/edit.tpl"
}

func (this *EditController) Post() {
	//数据处理
	this.Ct.Request.ParseForm()
	id := this.Ct.Request.Form.Get("id")
	atid, _ := strconv.Atoi(id)
	at := models.GetArticleById(atid)
	content := this.Ct.Request.Form.Get("content")
	at.Content = content
	models.InsertArticle(at)
	this.Ct.Redirect(302, "/admin/index")
}
