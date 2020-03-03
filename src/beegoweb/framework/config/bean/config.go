package bean

//全局通用配置类
var PropertyConfig Config

const (
	PRODUCTION string = "production"
	DEV        string = "dev"
	TEST       string = "test"
)

type Config struct {
	Env    string
	Name   string
	Server struct {
		Host string
		Port string
	}
	Redis struct {
		Url      string
		Password string
		Port     string
	}
	DataSource struct {
		Host         string
		DatabaseName string
		Username     string
		Password     string
		Port         string
		LogPath      string
	}
	LogConfig LogConfig
}
