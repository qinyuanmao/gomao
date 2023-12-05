package cfg

import (
	"fmt"
	"strings"

	"e.coding.net/tssoft/repository/gomao/db"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB(key string) (*db.MaoDB, error) {
	username := viper.GetString(fmt.Sprintf("%s.username", key))
	password := viper.GetString(fmt.Sprintf("%s.password", key))
	dbName := viper.GetString(fmt.Sprintf("%s.db_name", key))
	port := viper.GetString(fmt.Sprintf("%s.port", key))
	address := viper.GetString(fmt.Sprintf("%s.address", key))
	sslmode := viper.GetString(fmt.Sprintf("%s.sslmode", key))
	timezone := viper.GetString(fmt.Sprintf("%s.timezone", key))
	if sslmode == "" {
		sslmode = "disable"
	}
	if timezone == "" {
		timezone = "Europe/London"
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s TimeZone=%s", address, port, username, dbName, password, sslmode, timezone)
	var logMode = logger.Info
	if !viper.GetBool(fmt.Sprintf("%s.log_mode", key)) {
		logMode = logger.Silent
	}
	engine, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, err
	}
	return &db.MaoDB{DatabaseType: db.Postgres, DB: engine}, nil
}

func NewPostgresDBByENV(key string) (*db.MaoDB, error) {
	key = strings.ToUpper(key)
	username := viper.GetString(fmt.Sprintf("%s_USERNAME", key))
	password := viper.GetString(fmt.Sprintf("%s_PASSWORD", key))
	dbName := viper.GetString(fmt.Sprintf("%s_DB_NAME", key))
	port := viper.GetString(fmt.Sprintf("%s_PORT", key))
	address := viper.GetString(fmt.Sprintf("%s_ADDRESS", key))
	sslmode := viper.GetString(fmt.Sprintf("%s_SSLMODE", key))
	timezone := viper.GetString(fmt.Sprintf("%s_TIMEZONE", key))
	if sslmode == "" {
		sslmode = "disable"
	}
	if timezone == "" {
		timezone = "Europe/London"
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s TimeZone=%s", address, port, username, dbName, password, sslmode, timezone)
	var logMode = logger.Info
	if !viper.GetBool(fmt.Sprintf("%s.log_mode", key)) {
		logMode = logger.Silent
	}
	engine, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, err
	}
	return &db.MaoDB{DatabaseType: db.Postgres, DB: engine}, nil
}
