package controllers

import (
	"baoquan_ruanda/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}
//直接访问
func (l *LoginController)  Get() {
	//设置login_baoquan.html为模板文件tpl:template
	l.TplName="login_baoquan.html"

}
//用户登录接口
func (l *LoginController) Post() {
	var user models.Users
	err:=l.ParseForm(&user)
	if err!=nil {

		l.Ctx.WriteString("抱歉，用户信息解析失败")
		return
	}
	u,err:=user.QueryUser()
	if err!=nil {
		fmt.Println(err)
		l.Ctx.WriteString("抱歉,用户登录失败,请重试!")
		return
	}
	//登陆成功,跳转项目核心功能页面(storage.html)
	l.Data["name"]=u.Name
	l.TplName="storage.html"
}