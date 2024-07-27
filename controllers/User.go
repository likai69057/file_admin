package controllers

import (
	"file_admin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 注册页面控制器
type RegController struct {
	beego.Controller
}

func (this *RegController) ShowReg() {
	this.TplName = "register.html"
}

func (this *RegController) HandleReg() {
	//1.拿到浏览器传递的数据
	//通过GetString()函数 参数为获取的字段名称
	name := this.GetString("username")
	passwd := this.GetString("password")
	beego.Info(name, passwd)

	//2.数据处理
	if name == "" || passwd == "" {
		beego.Info("用户名或密码不能为空！")
		this.TplName = "register.html"
		return
	}

	//3.插入数据库
	//3.1获取orm对象
	o := orm.NewOrm()
	//3.2获取插入对象
	user := models.User{}
	//3.3插入操作
	user.UserName = name
	user.Passwd = passwd

	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("数据插入失败", err)
		return
	}
	//4.返回登录
	//重定向为登录页面
	this.Redirect("/", 302)
}

// 登录页面控制器
type LoginController struct {
	beego.Controller
}

func (this *LoginController) ShowLogin() {
	this.TplName = "login.html"
}

func (this *LoginController) HandleLogin() {
	//1.拿到浏览器传递的数据
	//通过GetString()函数 参数为获取的字段名称
	name := this.GetString("username")
	passwd := this.GetString("password")
	beego.Info(name, passwd)

	if name == "" || passwd == "" {
		beego.Info("用户名或密码不能为空！")
		this.TplName = "login.html"
	}

	//2.查询数据库
	//2.1获取orm对象
	o := orm.NewOrm()
	//2.2获取查询对象
	user := models.User{}
	//2.3查询操作
	user.UserName = name
	err := o.Read(&user, "UserName")
	if err != nil {
		beego.Info("用户名错误", err)
		this.TplName = "login.html"
		return
	}
	if passwd != user.Passwd {
		beego.Info("密码错误！")
		this.TplName = "login.html"
		return
	}

	//登录成功 跳转首页
	this.Redirect("/showArticle", 302)
}
