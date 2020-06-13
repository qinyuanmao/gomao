package cfg

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/qinyuanmao/gomao/db"
	"github.com/spf13/viper"
)

func NewPostgresDB(key string) (*db.MaoDB, error) {
	username := viper.GetString(fmt.Sprintf("%s.username", key))
	password := viper.GetString(fmt.Sprintf("%s.password", key))
	dbName := viper.GetString(fmt.Sprintf("%s.db_name", key))
	port := viper.GetString(fmt.Sprintf("%s.port", key))
	address := viper.GetString(fmt.Sprintf("%s.address", key))
	sslMode := viper.GetString(fmt.Sprintf("%s.sslmode", key))
	engine, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", address, port, username, dbName, password, sslMode))
	if err != nil {
		return nil, err
	}
	engine.LogMode(viper.GetBool(fmt.Sprintf("%s.log_mode", key)))
	return &db.MaoDB{db.Postgres, engine}, nil
}
