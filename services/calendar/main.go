package main

import (
	"calendar/app"
	"calendar/core/cache"
	"calendar/core/db"
	"calendar/core/opt"
	"calendar/core/redis"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	err := opt.ParseConfig("./conf/config.json")
	if err != nil {
		log.Fatalln("Parse Config err: ", err)
	}
	cfg := opt.Config()

	err = app.InitRouter()
	if err != nil {
		log.Fatalln("init router error: ", err)
	}

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

	ln, e := net.Listen("tcp", cfg.ServerAddr)
	log.Println("listen addr: ", cfg.ServerAddr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(ln, nil)
}
