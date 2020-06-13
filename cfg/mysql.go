package cfg

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/qinyuanmao/gomao/db"
	"github.com/spf13/viper"
)

func NewMysqlDB(key string) (*db.MaoDB, error) {
	username := viper.GetString(fmt.Sprintf("%s.username", key))
	password := viper.GetString(fmt.Sprintf("%s.password", key))
	dbName := viper.GetString(fmt.Sprintf("%s.db_name", key))
	address := viper.GetString(fmt.Sprintf("%s.address", key))
	engine, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", username, password, address, dbName))
	if err != nil {
		return nil, err
	}
	engine.LogMode(viper.GetBool(fmt.Sprintf("%s.log_mode", key)))
	return &db.MaoDB{db.Mysql, engine}, nil
}
