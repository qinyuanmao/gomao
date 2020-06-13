package cfg

import (
	"github.com/qinyuanmao/gomao/logger"
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
	return
}
