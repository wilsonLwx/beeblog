package main

import (
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"beeblog/models"
	"github.com/astaxie/beego/orm"
	"beeblog/controllers"
)

func init() {
	models.RegisterDB()
}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
	// 注册beego路由
	beego.Router("/",&controllers.MainController{})
	beego.Router("/login",&controllers.LoginController{})
	// 启动beego
	beego.Run()
}
