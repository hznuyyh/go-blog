package main

import (
	"github.com/astaxie/beego"
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	time.LoadLocation("Local")
	Database()
}

/**
 * 初始化数据库连接
 */
func Database() {
	//连接数据
	database_host := beego.AppConfig.String("DB_HOST")
	database_name := beego.AppConfig.String("DB_DATABASE")
	database_port := beego.AppConfig.String("DB_PORT")
	database_user := beego.AppConfig.String("DB_USERNAME")
	database_pwd := beego.AppConfig.String("DB_PASSWORD")
	DatabaseConnect(database_host,database_name,database_port,database_user,database_pwd,false)
}

func DatabaseConnect(database_host string , database_name string , database_port string , database_user string , database_pwd string , debug bool) {

	//选择模式
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = debug
	orm.DefaultTimeLoc = time.Local
	conn := database_user + ":" + database_pwd + "@tcp(" + database_host + ":" + database_port + ")/" + database_name + "?charset=utf8&parseTime=true&loc=Local"
	fmt.Println(conn)
	//注册数据库连接
	err := orm.RegisterDataBase("default", "mysql", conn)
	if err != nil {
		fmt.Println(err.Error())
	}
	//输出结果
	fmt.Printf("数据库连接成功！%s\n", conn)
}
