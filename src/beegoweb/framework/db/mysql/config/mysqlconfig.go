package database

import (
	"beegoweb/framework/config"
	"beegoweb/framework/config/bean"
	"fmt"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2"
	"sync/atomic"
)

var (
	engin *gorose.Engin
	dsc   = &bean.PropertyConfig.DataSource
	once  atomic.Value
)

func Register() {
	fmt.Println("声明式注册bean")
}

//主动调用
func InitDB() {
	initOnce()
}

func init() {
	logs.Info("database   init execute  ")
	config.ExecuteInit(func() {
		logs.Info("database util receive msg that config load down ,start init ")
		initOnce()
	})
}

func initOnce() {
	done := once.Load()
	if done == nil || !done.(bool) {
		initDB()
		once.Store(true)
	}
}

func initDB() {
	//dsn 用户名:密码@tcp(IP:端口号)/数据库名称?charset=utf8
	var dsn = dsc.Username + ":" +
		dsc.Password +
		"@tcp(" + dsc.Host + ":" + dsc.Port + ")/" + dsc.DatabaseName +
		"?charset=utf8&parseTime=true&loc=Asia%2FShanghai"
	var err error
	logs.Info("初始化数据库连接...")
	engin, err = gorose.Open(&gorose.Config{Driver: "mysql", Dsn: dsn})
	if err != nil {
		logs.Error("数据库连接失败! ", err)
	} else {
		logs.Info("数据库连接初始化完毕")
	}

	//engin.GetLogger().SetLogPath(dsc.LogPath)
	// engin.GetLogger().OpenLog()
}

func DB() gorose.IOrm {
	return engin.NewOrm()
}
