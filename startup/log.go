package startup

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var IoTLog = logs.NewLogger(10000) // 创建一个日志记录器，参数为缓冲区的大小

func init() {
	// 设置配置文件
	jsonConfig := `{
        "filename" : "` + beego.AppConfig.String("LogsFolder") + `IoT.log",
        "maxlines" : 1000,      
        "maxsize"  : 10240      
    }`
	IoTLog.SetLogger("file", jsonConfig) // 设置日志记录方式：本地文件记录
	IoTLog.SetLevel(logs.LevelDebug)     // 设置日志写入缓冲区的等级
	IoTLog.EnableFuncCallDepth(true)     // 输出log时能显示输出文件名和行号（非必须）
}
