package routers

import (
	"testbeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})

    //自动注册路由
	beego.Include(&controllers.ObjectController{})
}
