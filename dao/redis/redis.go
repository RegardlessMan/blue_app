package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"web_app/settings"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // 密码
		DB:       cfg.DB,       // 数据库
		PoolSize: cfg.PoolSize, // 连接池大小
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return
	}
	return nil
}

func Close() {
	_ = rdb.Close()
}
