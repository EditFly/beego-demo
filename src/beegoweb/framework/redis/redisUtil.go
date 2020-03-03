package redis

import (
	"beegoweb/common/util/json"
	"beegoweb/framework/config"
	"beegoweb/framework/config/bean"
	"fmt"
	"github.com/astaxie/beego/logs"
	goredis "github.com/go-redis/redis"
	"sync/atomic"
	"time"
)

var (
	redisConfig = &bean.PropertyConfig.Redis
	redis       *goredis.Client
	once        atomic.Value
)

func Register() {
	fmt.Println("声明式注册bean")
}

//主动初始化
func InitRedis() {
	initOnce()
}

func init() {
	logs.Info("redisUtil   init execute  ")
	config.ExecuteInit(func() {
		logs.Info("redis util receive msg that config load down ,start init ")
		initOnce()
	})
}
func initOnce() {
	done := once.Load()
	if done == nil || !done.(bool) {
		initRedis()
		once.Store(true)
	}
}

func initRedis() *goredis.Client {
	client := goredis.NewClient(&goredis.Options{
		Addr:     redisConfig.Url + ":" + redisConfig.Port,
		Password: redisConfig.Password, // no password set
		DB:       0,                    // use default DB
	})
	pong, err := client.Ping().Result()
	defer func() {
		if err != nil {
			client.Close()
			logs.Error("redis连接失败! ", pong, err)
		} else {
			logs.Info("redis初始化完毕")
		}
	}()
	return client
}

func Set(key string, value interface{}, expire time.Duration) error {
	err := redis.Set(key, value, expire).Err()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) (interface{}, error) {
	val, err := redis.Get(key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, nil
		}
		logs.Info(val, err.Error())
		return nil, err
	}
	return val, nil
}

func GetObj(key string, clazz interface{}) error {
	val, err := redis.Get(key).Result()
	if err != nil {
		return err
	}
	json.Parse(val, clazz)
	return nil
}

func Del(key string) error {
	err := redis.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}
