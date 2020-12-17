// @APIVersion 1.0.0
// @Title app API
// @Description app接口文档
package routers

import (
	"beegodemo/controllers"
	"github.com/astaxie/beego"
)

func init() {

	var names []*beego.Namespace

    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello",&controllers.HelloController{})
    ns:=beego.NewNamespace("mmkj",

    beego.NSNamespace("/gettest",beego.NSInclude(&controllers.GetTestController{})),
    )

    names=append(names,ns)
	beego.AddNamespace(names...)

}
