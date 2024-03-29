package cfg

import (
	"fmt"
	"strings"

	"e.coding.net/tssoft/repository/gomao/db"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB(key string) (*db.MaoDB, error) {
	username := viper.GetString(fmt.Sprintf("%s.username", key))
	password := viper.GetString(fmt.Sprintf("%s.password", key))
	dbName := viper.GetString(fmt.Sprintf("%s.db_name", key))
	address := viper.GetString(fmt.Sprintf("%s.address", key))
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", username, password, address, dbName)
	engine, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if viper.GetBool(fmt.Sprintf("%s.log_mode", key)) {
		engine = engine.Debug()
	}
	return &db.MaoDB{DatabaseType: db.Mysql, DB: engine}, nil
}

func NewMysqlDBByENV(key string) (*db.MaoDB, error) {
	key = strings.ToUpper(key)
	username := viper.GetString(fmt.Sprintf("%s_USERNAME", key))
	password := viper.GetString(fmt.Sprintf("%s_PASSWORD", key))
	dbName := viper.GetString(fmt.Sprintf("%s_DB_NAME", key))
	address := viper.GetString(fmt.Sprintf("%s_ADDRESS", key))
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", username, password, address, dbName)
	engine, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if viper.GetBool(fmt.Sprintf("%s.log_mode", key)) {
		engine = engine.Debug()
	}
	return &db.MaoDB{DatabaseType: db.Mysql, DB: engine}, nil
}
