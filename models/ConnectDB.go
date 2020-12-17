package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

func init()  {

	//runMode:=beego.BConfig.RunMode
	driverName:="mysql"

	//注册数据库驱动
	orm.RegisterDriver(driverName,orm.DRMySQL)

	//数据库连接
	dbUser:=beego.AppConfig.String("dbUser")
	dbPwd:=beego.AppConfig.String("dbPwd")
	dbHost:=beego.AppConfig.String("dbHost")
	dbPort:=beego.AppConfig.String("dbPort")
	dbName:=beego.AppConfig.String("dbName")

	dbConn:=dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8"

	err := orm.RegisterDataBase("default", driverName, dbConn)
	if err != nil {
		fmt.Println("连接数据库错误")
		return
	}
	fmt.Println("连接数据库成功")

}
func ConnectDB() {
	fmt.Println("----------------")

}
