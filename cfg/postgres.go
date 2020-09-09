package cfg

import (
	"fmt"

	"github.com/qinyuanmao/gomao/db"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(key string) (*db.MaoDB, error) {
	username := viper.GetString(fmt.Sprintf("%s.username", key))
	password := viper.GetString(fmt.Sprintf("%s.password", key))
	dbName := viper.GetString(fmt.Sprintf("%s.db_name", key))
	port := viper.GetString(fmt.Sprintf("%s.port", key))
	address := viper.GetString(fmt.Sprintf("%s.address", key))
	sslmode := viper.GetString(fmt.Sprintf("%s.sslmode", key))
	if sslmode == "" {
		sslmode = "disable"
	}

	dsn := fmt.Sprintf("host='%s' port='%s' user='%s' dbname='%s' password='%s' sslmode='%s'", address, port, username, dbName, password, sslmode)
	engine, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if viper.GetBool(fmt.Sprintf("%s.log_mode", key)) {
		engine = engine.Debug()
	}
	return &db.MaoDB{db.Postgres, engine}, nil
}
