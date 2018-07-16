// @APIVersion 1.0.0
// @Title IoTProject API
// @Description 物联网项目接口文档
package routers

import (
	"IoTProject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/test",
			beego.NSInclude(
				&controllers.TestController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
