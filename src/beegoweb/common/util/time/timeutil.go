package timeutil

import (
	"errors"
	"strings"
	"time"
)

//选择时间单位
func SwitchTimeUnit(str string) (time.Duration, error) {
	if str == "" {
		return 0, errors.New("不能为空")
	}
	str = strings.ToLower(str)
	switch str {
	case "microsecond":
		return time.Microsecond, nil
	case "millisecond":
		return time.Millisecond, nil
	case "second":
		return time.Second, nil
	case "minute":
		return time.Minute, nil
	case "hour":
		return time.Hour, nil
	}
	return 0, errors.New("没有匹配[microsecond,millisecond,second,minute,hour]")
}

func TimeToString(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
func TimeDayFormat(time time.Time) string {
	return time.Format("2006-01-02")
}
