package models

import (
	"baoquan_ruanda/db_baoquan"
	"beego_damo02/db_mysql"
	"crypto/md5"
	"encoding/hex"
)

type Users struct {
	Id int `form:"id"`
	Name string`form:"name"`
	Password string `form:"password"`
	//Myfile string	`input:"myfile"`
}
//保存用户信息到数据库里：使用方法
func (u Users) SaveUser() (int64,error) {
	//密码脱敏处理
	md5Hash:=md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes:=md5Hash.Sum(nil)
	u.Password=hex.EncodeToString(passwordBytes)





	//执行数据库操作
	row,err:=db_mysql.Db.Exec("insert into baoquan_registered (name,password)"+
		"values(?,?)",u.Name,u.Password)
	if err!=nil {
		return -1,err
	}
	id,err:=row.RowsAffected()
	if err!=nil {
		return -1,err
	}
	return id,nil
}
//
//查询用户信息
func (u Users) QueryUser()(*Users,error)  {
	md5Hash:=md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes:=md5Hash.Sum(nil)
	u.Password=hex.EncodeToString(passwordBytes)
	 row:=db_baoquan.Db.QueryRow("select name from baoquan_registered where name= ? and password= ?",
		u.Name,u.Password)
	//var name string
	err:= row.Scan(&u.Name)
	if err!=nil  {
		return nil,err
	}
	return &u,nil
}