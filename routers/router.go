package routers

import (
	"baoquan_ruanda/controllers"
	"github.com/astaxie/beego"
)
//路由
func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user_register",&controllers.RegisterController{})
    beego.Router("/login_baoquan.html",&controllers.LoginController{})
   beego.Router("/upload",&controllers.FileUploadController{})
    beego.Router("/upload_file.html",&controllers.FileUploadController{})
	beego.Router("/cert_detail.html",&controllers.CertDetailController{})
}
