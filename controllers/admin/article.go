package admin

import (
	"beepkg/models"
	"github.com/astaxie/beego"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Prepare() {

}

func (this *ArticleController) Get() {
	this.Ct.Request.ParseForm()
	page, _ := strconv.Atoi(this.Ct.Request.Form.Get("page"))
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
