package routers

import (
	"baoquan_ruanda/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user_register",&controllers.RegisterController{})
    beego.Router("/login_baoquan.html",&controllers.LoginController{})
   // beego.Router("/use_storage",&controllers.StorageController{})
}
