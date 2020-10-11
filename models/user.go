package models

import (
	"beego_damo02/db_mysql"
	"crypto/md5"
	"encoding/hex"
)

type Users struct {
	Id int `form:"id`
	Name string`form:"name"`
	Password string `form:"password"`
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