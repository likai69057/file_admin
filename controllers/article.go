package controllers

import (
	"file_admin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"path"
	"time"
)

// 注册页面控制器
type ArticleController struct {
	beego.Controller
}

// 跳转首页
func (this *ArticleController) ShowArticleList() {
	//1.查询数据库
	o := orm.NewOrm()
	//获取整个表结果queryseter对象
	qs := o.QueryTable("article")
	var articleResult []models.Article
	qs.All(&articleResult) // select * from article

	//2.将记录传送给视图
	this.Data["articleResult"] = articleResult

	this.TplName = "index.html"
}

// 跳转新增文章页面
func (this *ArticleController) ShowAddArticle() {
	this.TplName = "add.html"
}

// 上传文件
func (this *ArticleController) HandleAddArticle() {
	//1。拿数据
	articleName := this.GetString("article_name")
	articleContent := this.GetString("article_content")
	file, head, err := this.GetFile("upload_file")
	defer file.Close()

	//2.上传文件处理
	//1.判断上传文件的格式
	ext := path.Ext(head.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		beego.Info("上传文件格式不正确!")
		return
	}
	//2.判断文件大小
	if head.Size > 5000000 {
		beego.Info("文件太大，不允许上传!")
		return
	}

	//3.文件不能重名
	fileName := time.Now().Format("2006-01-02 15:04:05")

	this.SaveToFile("upload_file", "/static/img/"+fileName+ext)
	if err != nil {
		beego.Info("上传文件失败", err)
		return
	}

	//3.插入数据
	//1.获取orm对象
	o := orm.NewOrm()
	//2.获取数据对象
	articleOjb := models.Article{}
	//3.插入数据对象
	articleOjb.Title = articleName
	articleOjb.Content = articleContent
	articleOjb.Img = "/static/img/" + fileName + ext
	_, err = o.Insert(&articleOjb)
	if err != nil {
		beego.Info("插入数据失败", err)
	}

	//跳转列表页
	this.Redirect("/showArticle", 302)
}

// 跳转文章详情页
func (this *ArticleController) ShowArticleContent() {
	idsStr, err := this.GetInt("id")
	if err != nil {
		beego.Info("文章详情查询失败", err)
		return
	}

	o := orm.NewOrm()
	article := models.Article{}
	article.Id = idsStr

	err = o.Read(&article, "Id")
	if err != nil {
		beego.Info("没有该文章！", err)
		return
	}
	article.Count += 1
	o.Update(&article)
	this.Data["article"] = article

	this.TplName = "content.html"
}

// 文章删除
func (this *ArticleController) DeleteArticle() {
	idsStr, err := this.GetInt("id")
	if err != nil {
		beego.Info("文章详情查询失败", err)
		return
	}

	o := orm.NewOrm()
	article := models.Article{}
	article.Id = idsStr
	o.Delete(&article)
	this.Redirect("/showArticle", 302)
}
