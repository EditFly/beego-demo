package log

import (
	beanutil "beegoweb/common/util/bean"
	jsonutil "beegoweb/common/util/json"
	"beegoweb/framework/config/bean"
	notify2 "beegoweb/framework/notify"
	"github.com/astaxie/beego/logs"
)

var Log *logs.BeeLogger
var logConfig = &bean.PropertyConfig.LogConfig

func init() {
	Log = logs.GetBeeLogger()
	//Log.Reset()
	go waitConfig()
}

func waitConfig() {
	ok := <-notify2.LogConfigOK
	if ok {
		close(notify2.LogConfigOK)
		_init()
		notify2.LogLoadOk <- true
	}
}

func _init() {
	//异步输出
	Log.Async()
	//输出行号
	Log.EnableFuncCallDepth(true)
	confJson, err := jsonutil.Stringify(logConfig)
	if err != nil {
		logs.Error("日志配置文件转换失败！", err)
		return
	}
	mapdata := make(map[string]interface{})
	data := &mapdata
	err2 := jsonutil.Parse(confJson, data)
	if err2 != nil {
		logs.Error("日志配置文件转换失败！", err2)
		return
	}
	//剔除空属性
	beanutil.MapRmNil(mapdata)
	confJson, err = jsonutil.Stringify(mapdata)
	if err != nil {
		logs.Error("日志配置文件转换失败！", err)
		return
	}
	Log.SetLogger(logs.AdapterMultiFile, confJson)
	//Log.Warn("ddd")
}
