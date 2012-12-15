package admin

import (
	"beepkg/models"
	"github.com/astaxie/beego"
	"strconv"
)

type DelController struct {
	beego.Controller
}

func (this *DelController) Prepare() {

}

func (this *DelController) Get() {
	var at models.Article
	atId, _ := strconv.Atoi(this.Ct.Params[":id"])
	at = models.GetArticleById(atId)
	models.DeleteArticle(at)
	var pk models.PkgEntity
	pk.Id = at.Pkgid
	models.DeletePkg(pk)
	this.Ct.Redirect(302, "/admin/index")
}
