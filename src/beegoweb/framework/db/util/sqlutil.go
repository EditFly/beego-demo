package sqlutil

import stringUtil "beegoweb/common/util/string"

//批量插入
func BatchInsertSql(table string, fields []string, param [][]int) string {
	sql := "insert into " + table + "("
	lenf := len(fields)
	for i := 0; i < lenf; i++ {
		sql += fields[i]
		if i < lenf-1 {
			sql += ","
		}
	}
	sql += ") values "
	leni := len(param)
	for i := 0; i < leni; i++ {
		item := param[i]
		lenj := len(item)
		var _param = "("
		for j := 0; j < lenj; j++ {
			if v, e := stringUtil.ToString(item[j]); e == nil {
				_param += v
			} else {
				continue
			}
			if j != lenj-1 {
				_param += ","
			}
		}
		_param += ")"
		if i != leni-1 {
			_param += ","
		}
		sql += _param
	}
	return sql
}
