package main

import (
	"baoquan_ruanda/db_baoquan"
	_ "baoquan_ruanda/routers"
	"github.com/astaxie/beego"
)

func main() {


	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/login_js","./static/login_js")
	beego.SetStaticPath("/login_css","./static/login_css")
	beego.SetStaticPath("/login_img","./static/login_img")
	beego.SetStaticPath("/use_js","./static/use_js")
	beego.SetStaticPath("/use_css","./static/use_css")
	beego.SetStaticPath("/use_img","./static/use_img")
	db_baoquan.Init()
	beego.Run()
}

