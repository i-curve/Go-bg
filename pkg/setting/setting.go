package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	PageSize  int
	JwtSecret string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     time.Duration
	MaxActive   time.Duration
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("load the app.ini file error: \"%v\"", err)
	}
	MapTo("app", AppSetting)
	MapTo("server", ServerSetting)
	MapTo("database", DatabaseSetting)
	MapTo("redis", RedisSetting)
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}
func MapTo(name string, v interface{}) {
	err := cfg.Section(name).MapTo(v)
	if err != nil {
		log.Fatalf("load config file error: '%v'", err)
	}
}
