package service

import (
	"beegoweb/app/service/schedule/collection"
	database "beegoweb/framework/db/mysql/config"
	"beegoweb/framework/redis"
)

func DemoService() {
	redisAction()
}
func redisAction() {
	redis.Get("dd")
	db := database.DB()
	db.Get()

	collection.CollectDownload()
}
