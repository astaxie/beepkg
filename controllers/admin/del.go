package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beepkg/models"
	"strconv"
)

type DelController struct {
	beego.Controller
}

func (this *DelController) Prepare() {

}

func (this *DelController) Get() {
	var at models.Article
	atId, _ := strconv.Atoi(this.Ctx.Params[":id"])
	at = models.GetArticleById(atId)
	models.DeleteArticle(at)
	var pk models.PkgEntity
	pk.Id = at.Pkgid
	models.DeletePkg(pk)
	this.Ctx.Redirect(302, "/admin/index")
}
