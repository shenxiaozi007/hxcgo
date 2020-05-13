package main

import (
	"calendar/app"
	"calendar/core/cache"
	"calendar/core/opt"
	"calendar/core/redis"
	"calendar/core/rpc"
	"calendar/core/session"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	err := opt.ParseConfig("./conf/config.json")
	if err != nil {
		log.Fatalln("Parse Config err: ", err)
	}
	cfg := opt.Config()

	engine := gin.Default()
	engine.Static("/resource", cfg.ResourceDir)
	engine.StaticFS("/images", http.Dir(cfg.ImageDir))

	err = redis.Connect(cfg.Redis)
	if err != nil {
		log.Fatalln("redis error: ", err)
	}

	//注册缓存组件
	cacheDriver := cache.NewRedisDriver(redis.Client())
	cache.RegisterDriver(cacheDriver)

	err = rpc.Connect(cfg.RPC)
	if err != nil {
		log.Fatalln("rpc connect error: ", err)
	}

	//store, err := core.NewSessionStore(cfg.Session)
	//engine.Use(sessions.Sessions("session", store))

	engine.Use(session.New("token"))

	app.InitRouter(engine)
	engine.Run(cfg.ServerAddr)
}
