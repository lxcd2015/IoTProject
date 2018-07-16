package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["IoTProject/controllers:TestController"] = append(beego.GlobalControllerRouter["IoTProject/controllers:TestController"],
		beego.ControllerComments{
			Method: "GetData",
			Router: `/GetData`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
