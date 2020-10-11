package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}
//直接访问
func (l *LoginController)  Get() {
	//设置login_baoquan.html为模板文件tpl:template
	l.TplName="login_baoquan.html"

}