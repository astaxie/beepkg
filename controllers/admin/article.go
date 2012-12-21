package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beepkg/models"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Prepare() {

}

func (this *ArticleController) Get() {
	this.Ctx.Request.ParseForm()
	page, _ := strconv.Atoi(this.Ctx.Request.Form.Get("page"))
	offest := 400
	if page == 0 {
		page = 1
	}
	start := (page - 1) * offest
	this.Data["Articles"] = models.GetAllArticle(offest, start)
	//totalnums := models.GetTotalCount()
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/list.tpl"
}
