package conf

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type Config struct {
	Title  string
	Server ServerConfig `toml:"server"`
	Admin  AdminConfig  `toml:"admin"`
	Mysql  MysqlConfig  `toml:"mysql"`
	Redis  RedisConfig  `toml:"redis"`
}

type ServerConfig struct {
	IP             string `toml:"ip"`
	Port           string    `toml:"port"`
	PprofPort      int    `toml:"pprof_port"`
	AccessLog      string `toml:"access_log"`
	AccessLogLevel string `toml:"access_log_level"`
}

type AdminConfig struct {
	IP   string `toml:"ip"`
	Port string    `toml:"port"`
}

type MysqlConfig struct {
	DriverName      string `toml:"driver_name"`
	User            string `toml:"user"`
	PassWord        string `toml:"password"`
	IP              string `toml:"ip"`
	Port            int    `toml:"port"`
	DataBase        string `toml:"database"`
	MaxOpenConn     int    `toml:"max_open_conn"`
	MaxIdleConn     int    `toml:"max_idle_conn"`
	MaxConnLifeTime int    `toml:"max_conn_life_time"`
}

type RedisConfig struct {
	IP        string `toml:"ip"`
	Port      int    `toml:"port"`
	MaxActive int    `toml:"max_active"`
	MaxIdle   int    `toml:"max_idle"`
	DownGrade bool   `toml:"down_grade"`
}

var GlobalConfig Config

func init() {
	var config Config
	if _, err := toml.DecodeFile("./env/conf/base.toml", &config); err != nil {
		log.Panic("decode config file failed: %s", err)
		os.Exit(0)
	}
	GlobalConfig = config
	log.Println(config)
}
