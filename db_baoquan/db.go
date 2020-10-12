package db_baoquan

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
func Init(){
	fmt.Println("连接mysql数据库")
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")

	connUrl := dbUser +":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	fmt.Println(connUrl)
	db, err := sql.Open(dbDriver,connUrl)
	if err != nil {
		panic("数据库连接错误，请检查配置")
	}
	//为全局变量赋值
	Db = db
}
