package bean

type LogConfig struct {
	Filename string   `json:"filename"` //保存的文件名
	Maxlines int      `json:"maxlines"` //每个文件保存的最大行数，默认值 1000000
	Maxsize  int      `json:"maxsize"`  //每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
	Daily    bool     `json:"daily"`    //是否按照每天 logrotate，默认是 true
	Maxdays  int      `json:"maxdays"`  //文件最多保存多少天，默认保存 7 天
	Rotate   bool     `json:"rotate"`   //是否开启 logrotate，默认是 true
	Level    int      `json:"level"`    //日志保存的时候的级别，默认是 Trace 级别
	Perm     string   `json:"perm"`     //日志文件权限
	Separate []string `json:"separate"` //需要单独写入文件的日志级别,设置后命名类似 test.error.log
}
