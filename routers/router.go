package routers

import (
	"go-shiba/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
}
