package cfg

import (
	"e.coding.net/tssoft/repository/gomao/logger"
	"github.com/spf13/viper"
)

func LoadConfig(paths ...string) (err error) {
	viper.SetConfigType("yaml")
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	err = viper.ReadInConfig()
	if err != nil {
		logger.Error(err)
		return
	}
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	return
}
