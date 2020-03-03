package stringUtil

import (
	"errors"
	"strconv"
)

func SubStr(str string, start, end int) string {
	var substring = ""
	var pos = 0
	for _, c := range str {
		if pos < start {
			pos++
			continue
		}
		if pos >= end {
			break
		}
		pos++
		substring += string(c)
	}
	return substring
}

func ToString(val interface{}) (string, error) {
	switch val.(type) {
	case nil:
		return "", errors.New("没有找类型不能为nil")
	case int64:
		return strconv.FormatInt(val.(int64), 10), nil
	case int:
		return strconv.Itoa(val.(int)), nil
	case string:
		return val.(string), nil
	}
	return "", errors.New("没有找到匹配的数据类型")
}
