package models

import (
	"database/sql"
	"github.com/astaxie/beedb"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"strings"
)

type PkgEntity struct {
	Id       int    `PK` //主键
	Pid      int    //父ID
	Pathname string //包名，包括前缀路径 bufio、encoding/json、net/http
	Intro    string //包的简要介绍
}

type Article struct {
	Id      int `PK`
	Pkgid   int
	Content string
}

func GetAllPkg() (pkglist []PkgEntity) {
	db := GetLink()
	db.FindAll(&pkglist)
	return
}

func InsertPkg(pkg PkgEntity) PkgEntity {
	db := GetLink()
	db.Save(&pkg)
	return pkg
}

func DeletePkg(pkg PkgEntity) {
	db := GetLink()
	db.Delete(&pkg)
	return
}

func GetCruPkg(pkgname string) (crupkg PkgEntity) {
	db := GetLink()
	db.Where("pathname=?", strings.TrimRight(pkgname, "/")).Find(&crupkg)
	return
}

func GetPkgByID(id int) (crupkg PkgEntity) {
	db := GetLink()
	db.Where("id=?", id).Find(&crupkg)
	return
}

///获取pkg信息

func GetArticle(pkgid int) (article Article) {
	db := GetLink()
	db.Where("pkgid=?", pkgid).Find(&article)
	return
}

func GetArticleById(id int) (article Article) {
	db := GetLink()
	db.Where("id=?", id).Find(&article)
	return
}

func GetAllArticle(start, offset int) (atlist []Article) {
	db := GetLink()
	db.Limit(start, offset).FindAll(&atlist)
	return
}

func GetTotalCount() int {
	db := GetLink()
	a, _ := db.SetTable("article").SetPK("id").Select("count(*) as num").FindMap()
	var num int
	if len(a) >= 1 {
		num, _ = strconv.Atoi(string(a[0]["num"]))
	} else {
		num = 0
	}
	return num
}

func InsertArticle(at Article) Article {
	db := GetLink()
	db.Save(&at)
	return at
}

func DeleteArticle(at Article) {
	db := GetLink()
	db.Delete(&at)
	return
}

func UpdateArticle(at Article) Article {
	db := GetLink()
	db.Save(&at)
	return at
}

func GetLink() beedb.Model {
	db, err := sql.Open("sqlite3", "./beepkg.db")
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)
	return orm
}
