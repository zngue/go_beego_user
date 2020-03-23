package controllers

import "github.com/astaxie/beego"

type User struct {
	beego.Controller
}

func (u *User) Index()  {
	u.Data["json"]= map[string]interface{}{ "name":"zhangsan","age":24 }
	u.ServeJSON()
}
func (u *User) Add ()  {

	u.TplName="user/add.html"
}