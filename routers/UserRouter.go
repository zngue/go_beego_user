package routers

import (
	"github.com/astaxie/beego"
	"user_model/controllers"
)
func UserRouter()  {
	ns := beego.NewNamespace("/user",
		beego.NSRouter("/index",&controllers.User{},"*:Index"),
		beego.NSRouter("add",&controllers.User{},"*:Add"),
	)
	beego.AddNamespace(ns)
}