package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id       int
	UserName string
	Passwd   string
}

// 表的设计
type Article struct {
	Id      int       `orm:"pk;auto"`
	Title   string    `orm:"size(20)"`                    //文章标题
	Content string    `orm:"size(500)"`                   //文章内容
	Img     string    `orm:"size(50);null"`               //图片
	Time    time.Time `orm:"type(datetime);auto_now_add"` //发布时间
	Count   int       `orm:"default(0)"`                  //阅读量
}

func init() {
	//1.连接数据库
	link := "root:Alikai19940818@tcp(114.132.222.214:3306)/newsWeb?charset=utf8"
	orm.RegisterDataBase("default", "mysql", link)
	orm.RegisterModel(new(User), new(Article))
	orm.RunSyncdb("default", false, true)
}
