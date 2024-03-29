package cfg

import (
	"fmt"
	"strings"
	"time"

	"e.coding.net/tssoft/repository/gomao/logger"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

type RedisDB struct {
	*redis.Client
}

func NewRedisConf(key string) (*RedisDB, error) {
	redisOptions := &redis.Options{
		Password: viper.GetString(fmt.Sprintf("%s.password", key)),
		DB:       viper.GetInt(fmt.Sprintf("%s.db", key)),
		Addr:     viper.GetString(fmt.Sprintf("%s.address", key)),
	}
	redisClient := redis.NewClient(redisOptions)
	if pong, err := redisClient.Ping().Result(); err != nil {
		logger.Error(err.Error())
		return nil, err
	} else {
		logger.Info(pong)
	}
	return &RedisDB{redisClient}, nil
}

func NewRedisConfByENV(key string) (*RedisDB, error) {
	key = strings.ToUpper(key)
	redisOptions := &redis.Options{
		Password: viper.GetString(fmt.Sprintf("%s_PASSWORD", key)),
		DB:       viper.GetInt(fmt.Sprintf("%s_DB", key)),
		Addr:     viper.GetString(fmt.Sprintf("%s_ADDRESS", key)),
	}
	redisClient := redis.NewClient(redisOptions)
	if pong, err := redisClient.Ping().Result(); err != nil {
		logger.Error(err.Error())
		return nil, err
	} else {
		logger.Info(pong)
	}
	return &RedisDB{redisClient}, nil
}

func (rdb *RedisDB) Save(key string, value interface{}, expiredAt time.Duration) (err error) {
	if rdb.IsExist(key) {
		if err = rdb.Del(key).Err(); err != nil {
			return
		}
	}
	return rdb.Set(key, value, expiredAt).Err()
}

func (rdb *RedisDB) IsExist(key string) bool {
	return rdb.Exists(key).Val() > 0
}
