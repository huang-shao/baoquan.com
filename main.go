package main

import (
	"baoquan_ruanda/blockchain"
	"baoquan_ruanda/db_baoquan"
	"baoquan_ruanda/models"
	_ "baoquan_ruanda/routers"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	user1:=models.Users{
		Id:       1,
		Name:     "18770882156",
		Password: "123456",
	}

	fmt.Println(user1)

json.Unmarshal()



	return
	block:=blockchain.NewBlock(0,[]byte{},[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(block)
	fmt.Printf("区块hash值:%x\n",block.Hash)
	fmt.Printf("区块hash值:%d\n",block.Nonce)
	return
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

