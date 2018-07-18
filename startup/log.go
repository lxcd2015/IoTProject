package startup

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// var log = logs.NewLogger(10000) // 创建一个日志记录器，参数为缓冲区的大小
var IoTLog, IoTSystemLog *IoTLogHelp

func init() {

	IoTLog = new(IoTLogHelp)
	log := logs.NewLogger(10000)
	log.SetLevel(logs.LevelDebug) // 设置日志写入缓冲区的等级
	log.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）
	// 设置配置文件
	jsonConfig := `{
        "filename" : "` + beego.AppConfig.String("LogsFolder") + `IoT.log",
        "maxlines" : 1000,      
        "maxsize"  : 10240      
    }`
	log.SetLogger("file", jsonConfig) // 设置日志记录方式：本地文件记录
	IoTLog.log = log

	IoTSystemLog = new(IoTLogHelp)
	systemLog := logs.NewLogger(10000)
	systemLog.SetLevel(logs.LevelDebug) // 设置日志写入缓冲区的等级
	systemLog.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）
	// 设置配置文件
	jsonConfig = `{
			"filename" : "` + beego.AppConfig.String("LogsFolder") + `IoTSystem.log",
			"maxlines" : 1000,      
			"maxsize"  : 10240      
		}`
	systemLog.SetLogger("file", jsonConfig) // 设置日志记录方式：本地文件记录
	IoTSystemLog.log = systemLog
}

type IoTLogHelp struct {
	log *logs.BeeLogger
}

func (lg *IoTLogHelp) Emergency(format string, v ...interface{}) {
	lg.log.Emergency(format+"\r", v...)
}
func (lg *IoTLogHelp) Alert(format string, v ...interface{}) {
	lg.log.Alert(format+"\r", v...)
}
func (lg *IoTLogHelp) Critical(format string, v ...interface{}) {
	lg.log.Critical(format+"\r", v...)
}
func (lg *IoTLogHelp) Error(format string, v ...interface{}) {
	lg.log.Error(format+"\r", v...)
}
func (lg *IoTLogHelp) Warning(format string, v ...interface{}) {
	lg.log.Warning(format+"\r", v...)
}
func (lg *IoTLogHelp) Notice(format string, v ...interface{}) {
	lg.log.Notice(format+"\r", v...)
}
func (lg *IoTLogHelp) Informational(format string, v ...interface{}) {
	lg.log.Informational(format+"\r", v...)
}
func (lg *IoTLogHelp) Debug(format string, v ...interface{}) {
	lg.log.Debug(format+"\r", v...)
}
