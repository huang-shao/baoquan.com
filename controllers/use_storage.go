package controllers

import "github.com/astaxie/beego"

type StorageController struct {
	beego.Controller
}

func (s *StorageController) Get() {
	s.TplName="/storage.html"
}