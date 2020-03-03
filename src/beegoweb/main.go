package main

import (
	_ "beegoweb/app/routers"
	"beegoweb/framework/notify"
	"beegoweb/framework/start"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	//log.Println("等待各组件初始化完毕....")
	start.StartSystem()
	notify.MainWait.Wait()
	logs.Info("全部组件初始化完毕，启动服务")
	env := beego.BConfig.RunMode
	logs.Info("开始选择配置文件,当前环境：", env)
	beego.Run()
}
