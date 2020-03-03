package random

import (
	"github.com/astaxie/beego/logs"
	"math/rand"
	"sync"
)

type snowflake struct {
	worker *Worker
}

var sfw *snowflake
var once sync.Once

//单例模式
func instanceSnowflake() *snowflake {
	workerId := RandInt64(0, 1023)
	once.Do(func() {
		logs.Info("only one will do this function")
		worker, err := NewWorker(workerId)
		if err != nil {
			//todo
		}
		sfw = &snowflake{
			worker: worker,
		}
	})
	return sfw
}

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 && max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

//通过雪花算法获取唯一ID
func GetUUID() (int64, error) {
	if sfw == nil {
		sfw = instanceSnowflake()
	}
	return sfw.worker.GetSnowflakeId(), nil
}

//func GetUUIDByWid(workerId int64) (int64, error) {
//	worker, err := util.NewWorker(workerId)
//	if err != nil {
//		return 0, err
//	}
//	return worker.GetSnowflakeId(), nil
//}
