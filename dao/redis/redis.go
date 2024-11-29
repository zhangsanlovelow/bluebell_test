package redis

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client

// 初始化redis连接
func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		// Addr:     "localhost:6379",
		Addr:     viper.GetString("redis.addr") + ":" + viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.db"),          // use default DB
	})
	_, err = rdb.Ping().Result()
	return err

}

func Close() {
	_ = rdb.Close()
}
