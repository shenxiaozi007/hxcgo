package main

import (
	"github.com/huangxinchun/hxcgo/admin/app"
	"github.com/huangxinchun/hxcgo/admin/core/cache"
	"github.com/huangxinchun/hxcgo/admin/core/db"
	"github.com/huangxinchun/hxcgo/admin/core/opt"
	"github.com/huangxinchun/hxcgo/admin/core/redis"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	//获取当前目录
	pwd, _ := os.Getwd()

	err := opt.ParseConfig(pwd + "/conf/config.json")
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

	//注册缓存组件
	cacheDriver := cache.NewRedisDriver(redis.Client())
	cache.RegisterDriver(cacheDriver)

	log.Println("debug: ")
	log.Println(redis.Client().Set("test", "111", 10*time.Second).Result())

	ln, e := net.Listen("tcp", cfg.ServerAddr)
	log.Println("listen addr: ", cfg.ServerAddr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(ln, nil)
}
