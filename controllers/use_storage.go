package controllers

import (
	"github.com/astaxie/beego"
)

type StorageController struct {
	beego.Controller
}

func (s *StorageController) Post() {
	//fmt.Println("============================2345678=========================================")
	s.TplName="storage.html"
}