package config

import (
	"beegoweb/common/util/sys"
	"beegoweb/framework/config/bean"
	"beegoweb/framework/log"
	"beegoweb/framework/notify"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/spf13/viper"
	"os"
	"sync/atomic"
)

var notifyInitNum int32 = 0

const applicationName = ".application.yml"

func init() {
	beego.LoadAppConfig("ini", "/conf/app.conf")
}

func AutoChoiceEnvConfig() {
	env := beego.BConfig.RunMode
	log.Log.Info("开始选择配置文件,当前环境：", env)
	logs.Info("开始选择配置文件,当前环境：", env)
	filename := func(env string) string {
		switch env {
		//正式
		case bean.PRODUCTION:
			return "production"
			//测试
		case bean.DEV:
			return "dev"
			//本地
		case bean.TEST:
			return "test"
		}
		logs.Error("没有设置环境参数，加载配置文件失败！")
		return ""
	}(env)
	err := loadConfig(filename + applicationName)
	if err != nil {
		logs.Error(err)
		sys.CloseSystem()
	}
}

func loadConfig(filename string) error {
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		logs.Error(err)
		sys.CloseSystem()
	}
	path += "/conf/"
	configLoad := viper.New()
	configLoad.AddConfigPath(path)     //设置读取的文件路径
	configLoad.SetConfigName(filename) //设置读取的文件名
	configLoad.SetConfigType("yaml")   //设置文件的类型
	//尝试进行配置读取
	if err := configLoad.ReadInConfig(); err != nil {
		return err
	}
	//实例化
	unmarshalErr := configLoad.Unmarshal(&bean.PropertyConfig)
	if unmarshalErr != nil {
		logs.Error("实例化配置文件失败", unmarshalErr)
		sys.CloseSystem()
		return errors.New("实例化配置文件失败")
	}
	//通知配置文件加载完毕
	notify.MainWait.Add(1)
	notify.LogConfigOK <- true
	go waitLogLoading()
	return nil
}

func waitLogLoading() {
	ok := <-notify.LogLoadOk
	if ok {
		notify.MainWait.Done()
		close(notify.LogLoadOk)
		notifyInitFnc()
	}
}

func notifyInitFnc() {
	defer func() {
		logs.Info("关闭 config 通知管道")
		close(notify.LoadDownConfig)
	}()
	logs.Info("有", notifyInitNum, "个init注册,等待执行")
	for ; notifyInitNum > 0; notifyInitNum-- {
		notify.LoadDownConfig <- true
	}
}

//增加需要通知 config加载完毕的方法
func addInitNum() {
	notify.MainWait.Add(1)
	atomic.AddInt32(&notifyInitNum, 1)
}
func ExecuteInit(f func()) {
	addInitNum()
	go func() {
		//监听配置文件加载完毕信号
		ok := <-notify.LoadDownConfig
		if ok {
			f()
			//执行完毕后在
			notify.MainWait.Done()
		}
	}()
}
