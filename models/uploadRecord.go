package models

import (
	"baoquan_ruanda/db_baoquan"
	"baoquan_ruanda/util"
	"fmt"
)

/**
 *上传文件记录的 结构体定义
 */


type UploadRecord struct {
	Id int
	FileName string
	FileSize int64
	FileCert string//认证号
	FileTitle string
	CertTime int64
	FormatCertTime string
	User_Name string//对应的用户的name
}
/*
  保存上传记录到数据库中
 */
func (u UploadRecord) SaveRecord() (int64,error) {
	fmt.Println("认证数据的用户信息是：",u.User_Name)
	rs,err:=db_baoquan.Db.Exec("insert into upload_record(file_name,file_size,file_cert,file_title,cert_time,user_name)"+
		"values(?,?,?,?,?,?)",
		u.FileName,
		u.FileSize,
		u.FileCert,
		u.FileTitle,
		u.CertTime,
		u.User_Name)
	if err!=nil {
		return -1,err
	}
	id,err:=rs.RowsAffected()
	if err!=nil {
		return -1,err
	}
	return id,nil
}

/*
  读取数据库在用户对应的所有认证数据
 */
func QueryRecord(user_name1 string)([]UploadRecord,error)  {
	rs,err:=db_baoquan.Db.Query("select id,file_name,file_size,file_cert,file_title,cert_time,user_name from upload_record where user_name = ?",user_name1)
	if err!=nil {
		return nil,err
	}
	records:=make([]UploadRecord,0)
	for rs.Next()  {
		var record UploadRecord
		err:=rs.Scan(&record.Id,&record.FileName,&record.FileSize,&record.FileCert,&record.FileTitle,&record.CertTime,&record.User_Name)
		if err!=nil {
			return nil,err
		}
		record.FormatCertTime=util.TimeFormat(record.CertTime,0,util.TIME_FORMAT_FOUR)
		records=append(records,record)
	}
	return records,nil
}