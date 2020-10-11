package controllers

import (
	"baoquan_ruanda/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Get() {
	//解析请求数据
	var user models.Users
	err:=r.ParseForm(&user)
	if err!=nil{
		r.Ctx.WriteString("抱歉，解析错误!")
		return
	}
	//保存用户数据到数据库
	_,err=	user.SaveUser()
	if err!=nil {
		r.Ctx.WriteString("抱歉，用户注册失败！")
		return
	}
	//用户注册成功
	r.TplName="login_baoquan.html"

	//
	//bodyBytes,err :=ioutil.ReadAll(r.Ctx.Request.Body)
	//if err != nil {
	//	r.Ctx.WriteString("数据接收错误,请重试")
	//	return
	//}
	//var user models.Users
	//err = json.Unmarshal(bodyBytes,&user)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	r.Ctx.WriteString("数据解析错误")
	//	return
	//}
	//
	////3、将解析到的用户数据，保存到数据
	//id, err := db_baoquan.InsertUser(user)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	r.Ctx.WriteString("用户保存失败.")
	//	return
	//}
	//fmt.Println(id)
	//
	//r.ServeJSON()
}