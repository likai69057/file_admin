package main

import (
	_ "file_admin/models"
	_ "file_admin/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
