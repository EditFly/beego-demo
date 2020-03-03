package collection

import (
	stringUtil "beegoweb/common/util/string"
	timeutil "beegoweb/common/util/time"
	database "beegoweb/framework/db/mysql/config"
	"github.com/astaxie/beego/logs"
	"github.com/gohouse/gorose/v2"
	"strconv"
	"time"
)

const orgTbName = "down_detail"
const collectDownloadTbName = "collection_download"

func CollectDownload() {
	db := database.DB()
	_now := time.Now()
	now := timeutil.TimeDayFormat(_now)
	//查询下载次数

	downCount, err := db.Query("select count(0) as down_count,content_id, content_type, sub_id, type "+
		"from down_detail "+
		" where DATE_FORMAT(down_time,'%Y-%m-%d')= ? "+
		" group by content_id, content_type, sub_id, type ", now)
	if err != nil {
		logs.Error("执行查询失败", err)
		return
	}

	//查询下载人次并合并
	var sql1 = "SELECT  content_id,content_type,sub_id,type FROM down_detail " +
		" where DATE_FORMAT(down_time,'%Y-%m-%d') = ? " +
		" GROUP BY content_id, content_type, sub_id, type ,user_id "
	var sql2 = " SELECT p.*,count(0) as down_person_count from (" + sql1 + " ) p GROUP BY p.content_id,p.content_type,p.sub_id,p.type "
	personDown, err := db.Query(sql2, now)
	if err != nil {
		logs.Error("执行查询失败", err)
		return
	}
	//log.Println("downCount= ", downCount, "personDown= ", personDown)
	if downCount != nil && len(downCount) > 0 {
		len := len(downCount)
		var item gorose.Data
		for i := 0; i < len; i++ {
			item = downCount[i]
			//item["down_count"] = item["count"]
			inertAndUpdate(db, item)
		}
	}

	if personDown != nil && len(personDown) > 0 {
		len := len(personDown)
		var item gorose.Data
		for i := 0; i < len; i++ {
			item = personDown[i]
			inertAndUpdate(db, item)
		}
	}
}

func inertAndUpdate(db gorose.IOrm, item gorose.Data) {
	db.Reset()
	//查询已存在的记录
	rowdata, err := db.Table(collectDownloadTbName).Fields("id", "down_person_count", "down_count").
		Where("content_id", item["content_id"]).
		Where("content_type", item["content_type"]).
		Where("sub_id", item["sub_id"]).
		Where("type", item["type"]).First()
	if err != nil {
		logs.Error("执行查询失败", err)
	}
	db.Reset()
	if rowdata != nil { //更新

		var sql = "update collection_download "
		if item["down_count"] != nil {
			s, err := stringUtil.ToString(item["down_count"])
			if err != nil {
				logs.Error(err)
				return
			}
			sql += " set down_count=down_count+" + s
			//item["down_count"].(int)
		}
		if item["down_person_count"] != nil {
			s, err := stringUtil.ToString(item["down_person_count"])
			if err != nil {
				logs.Error(err)
				return
			}
			sql += "set down_person_count=down_person_count+" + s
		}
		sql += " where id=" + strconv.FormatInt(rowdata["id"].(int64), 10)
		db.Execute(sql)
	} else { //插入
		db.Table(collectDownloadTbName).Data(item).Insert()
		//batchInsert(collectDownloadTbName,[]{"",""},[]{}{})
	}
}
