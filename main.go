package main

import (
	"beegodemo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "beegodemo/routers"
)

func main() {

	//初始化参数传递
	params:=make(map[string]interface{})
	//初始化日志
	initLog(params)

	//数据库初始化
	initDataBase(params)

	//异步输出
	logs.Async()

	beego.Run()
}

func initLog(params map[string]interface{})  {
	//log 记录设置
	logs.SetLogFuncCall(true) //设置calldepth开启 ,调用栈上多少层的信息，默认为4
	logs.SetLogFuncCallDepth(3)

	//log保存设置

	//filename 保存文件名
	//maxLines 每个文件保存的最大行数，默认值1000000
	//maxsize 每个文件保存的最大尺寸 默认1<<28   256MB
	//daily 是否按照每天 logrotate 默认true
	// maxDays 文件最多保存多少天 ,默认保存7
	//rotate  是否开启 logrotate 默认true
	//level 日志保存的级别 ，默认trace
	//perm 日志文件权限
	_ = logs.SetLogger("file", `{"filename":"./logs/logs.log","perm":"0775","maxDays":15}`)

	//设置文件行号
	logs.EnableFuncCallDepth(true)

	//等级
	logs.SetLevel(logs.LevelDebug)
	beego.BConfig.WebConfig.DirectoryIndex=true //应用程序默认配置

	//自动文档
	beego.BConfig.WebConfig.StaticDir["/swagger"]="swagger"
	orm.Debug=true

}
//初始化连接
func initDataBase(params map[string]interface{})  {
	//数据库连接
	models.ConnectDB()
}