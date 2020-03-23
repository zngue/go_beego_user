package routers

import (
	"user_model/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    UserRouter()
}
