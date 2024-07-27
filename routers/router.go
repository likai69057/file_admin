package routers

import (
	"file_admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//登录页面
	beego.Router("/", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
	//注册页面
	beego.Router("/register", &controllers.RegController{}, "get:ShowReg;post:HandleReg")
	//首页
	beego.Router("/showArticle", &controllers.ArticleController{}, "get:ShowArticleList")
	//插入文章
	beego.Router("/addArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")
	//文章详情页
	beego.Router("/articleContent", &controllers.ArticleController{}, "get:ShowArticleContent")
	//删除文章
	beego.Router("/deleteArticle", &controllers.ArticleController{}, "get:DeleteArticle")
}
