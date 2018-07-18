package startup

import (
	"github.com/astaxie/beego"
	. "github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeStatic, beforeStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, beforeRouter)
	beego.InsertFilter("/*", beego.BeforeExec, beforeExec)
	beego.InsertFilter("/*", beego.AfterExec, afterExec)
	beego.InsertFilter("/*", beego.FinishRouter, finishRouter)
}

//静态地址之前
func beforeStatic(ctx *Context) {

}

//寻找路由之前
func beforeRouter(ctx *Context) {

}

//找到路由之后，开始执行相应的 Controller 之前
func beforeExec(ctx *Context) {
	writeOperationLog(ctx)
}

//执行完 Controller 逻辑之后执行的过滤器
func afterExec(ctx *Context) {

}

//执行完逻辑之后执行的过滤器
func finishRouter(ctx *Context) {

}

//记录操作日志
func writeOperationLog(ctx *Context) {
	IoTSystemLog.Informational("Uri：%v；方法：%v；IP地址：%v；UserAgent：%v；",  ctx.Input.URI(), ctx.Input.Method(),ctx.Input.IP(), ctx.Input.UserAgent())
}
