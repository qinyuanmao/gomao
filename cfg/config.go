package cfg

import (
	"e.coding.net/tssoft/repository/gomao/logger"
	"github.com/spf13/viper"
)

func LoadConfig(path string) (err error) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	err = viper.ReadInConfig()
	if err != nil {
		logger.Error(err)
		return
	}
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	return
}
