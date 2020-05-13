package api

import (
	"log"
	"testing"
	"time"
	"wechat/app/proto"
	"wechat/core/db"
	"wechat/core/opt"
	"wechat/core/redis"
	"wechat/core/cache"
)

func TestAccessToken(t *testing.T) {
	err := opt.ParseConfig("../../conf/config.json")
	if err != nil {
		log.Fatalln("Parse Config err: ", err)
	}
	cfg := opt.Config()

	err = db.Connect(cfg.DB)
	if err != nil {
		log.Fatalln("db connect error: ", err)
	}

	err = redis.Connect(cfg.Redis)
	if err != nil {
		log.Fatalln("redis connect error: ", err)
	}

	log.Println("debug: ")
	log.Println(redis.Client().Set("test", "111", 10*time.Second).Result())

	//注册缓存组件
	cacheDriver := cache.NewRedisDriver(redis.Client())
	cache.RegisterDriver(cacheDriver)

	accessToken,err := AccessToken("wx57442d9c0889a01c","3ebffc4eb45538d6724e6ca3041b2292")
	if err != nil {
		log.Fatalln(err)
	}

	args := &proto.QRCodeUnlimitedReq{
		Scene:       "test",
		Page:        "",
		Width:       0,
		AutoColor:   false,
		LineColor:   nil,
		IsHyaline:   false,
	}

	log.Println(AccessToken("wx57442d9c0889a01c","3ebffc4eb45538d6724e6ca3041b2292"))
	QRCodeUnlimited(accessToken,args)
}
