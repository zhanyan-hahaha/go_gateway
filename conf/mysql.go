package conf

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MySQLConf struct {
	DriverName      string `mapstructure:"driver_name"`
	DataSourceName  string `mapstructure:"data_source_name"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxConnLifeTime int    `mapstructure:"max_conn_life_time"`
}

type MysqlMapConf struct {
	List map[string]*MySQLConf `mapstructure:"list"`
}

//var DBMapPool map[string]*sql.DB
//var GORMMapPool map[string]*gorm.DB
//var DBDefaultPool *sql.DB
//var GORMDefaultPool *gorm.DB
var DB *gorm.DB

func InitDBPool(config Config) error {
	fmt.Println(config.Mysql.DataSourceName)
	db, err := gorm.Open(config.Mysql.DriverName, config.Mysql.DataSourceName)
	fmt.Println(db)
	if err != nil{
		return err
	}
	DB = db
	return nil
}


//func GetDBPool(name string) (*sql.DB, error) {
//	if dbpool, ok := DBMapPool[name]; ok {
//		return dbpool, nil
//	}
//	return nil, errors.New("get pool error")
//}
//
//func GetGormPool(name string) (*gorm.DB, error) {
//	if dbpool, ok := GORMMapPool[name]; ok {
//		return dbpool, nil
//	}
//	return nil, errors.New("get pool error")
//}