package controllers

import (
	. "IoTProject/startup"

	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

// @Title 获取数据
// @Description 获取测试用数据
// @Success 200
// @router /GetData [get]
func (u *TestController) GetData() {
	u.Data["json"] = "测试数据2018"
	IoTLog.Error("Error0718测试成功")
	// u.Abort("500")
	// u.Data["json"] = u.Data["json"] + "2018年7月15日，测试成功"
	u.ServeJSON()
}
