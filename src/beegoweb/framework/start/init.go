package start

import (
	"beegoweb/framework/config"
	database "beegoweb/framework/db/mysql/config"
	"beegoweb/framework/redis"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/config/yaml"
)

func StartSystem() {
	beego.BConfig.Log.AccessLogs = true
	//加载配置文件
	config.AutoChoiceEnvConfig()
	//手动初始化
	database.Register()
	redis.Register()
}
