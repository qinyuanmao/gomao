package cfg

import (
	"fmt"

	"e.coding.net/tssoft/repository/gomao/db"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

func NewClickHouseDB(key string) (*db.MaoDB, error) {
	username := viper.GetString(fmt.Sprintf("%s.username", key))
	password := viper.GetString(fmt.Sprintf("%s.password", key))
	dbName := viper.GetString(fmt.Sprintf("%s.db_name", key))
	address := viper.GetString(fmt.Sprintf("%s.address", key))
	writeTimeout := viper.GetInt(fmt.Sprintf("%s.write_timeout", key))
	readTimeout := viper.GetInt(fmt.Sprintf("%s.read_timeout", key))
	dsn := fmt.Sprintf("tcp//%s?database=%s&username=%spassword=%s&read_timeout=%d&write_timeout=%d", address, dbName, username, password, readTimeout, writeTimeout)
	engine, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if viper.GetBool(fmt.Sprintf("%s.log_mode", key)) {
		engine = engine.Debug()
	}
	return &db.MaoDB{DatabaseType: db.Clickhouse, DB: engine}, nil
}
