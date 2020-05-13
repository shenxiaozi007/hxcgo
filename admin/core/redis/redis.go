package redis

import (
	"github.com/huangxinchun/hxcgo/admin/core"
	"fmt"
	"github.com/go-redis/redis"
)

var clients = map[string]*redis.Client{}

func Connect(configs []*core.RedisConfig) error {
	var defaultClient *redis.Client

	for _,cfg := range configs {
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d",cfg.Host,cfg.Port),
			Password: cfg.Password, // no password set
			DB:       cfg.Database,  // use default DB
		})

		if err := client.Ping().Err();err != nil {
			return err
		}
		if defaultClient == nil || cfg.Alias == "default" {
			clients["default"] = client
		}
	}

	return nil
}

func Client(alias ...string) *redis.Client {
	var redisName string
	if len(alias) == 0 {
		redisName = "default"
	}else{
		redisName = alias[0]
	}

	client,ok := clients[redisName]
	if !ok {
		panic(fmt.Sprintf("redis %s does not exists",redisName))
	}

	return client
}
