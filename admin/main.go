package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/huangxinchun/hxcgo/admin/app"
	"github.com/huangxinchun/hxcgo/admin/core"
	"github.com/huangxinchun/hxcgo/admin/core/cache"
	"github.com/huangxinchun/hxcgo/admin/core/opt"
	"github.com/huangxinchun/hxcgo/admin/core/redis"
	"github.com/huangxinchun/hxcgo/admin/core/rootpath"
	"github.com/huangxinchun/hxcgo/admin/core/rpc"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	// err := opt.ParseConfig("/home/vagrant/gocode/hxcgo/admin/conf/config.json")
	err := opt.ViperConfig(path.Join(rootpath.RootPath, "/conf"))
	if err != nil {
		log.Fatalln("Parse Config err: ", err)
	}
	cfg := opt.Config()

	engine := gin.Default()

	err = redis.Connect(cfg.Redis)
	if err != nil {
		log.Fatalln("redis error: ", err)
	}
	log.Println("debug: ")
	log.Println(redis.Client().Set("test", "111", 10*time.Second).Result())

	//注册缓存组件
	cacheDriver := cache.NewRedisDriver(redis.Client())
	cache.RegisterDriver(cacheDriver)

	err = rpc.Connect(cfg.RPC)
	if err != nil {
		log.Fatalln("rpc connect error: ", err)
	}

	store, err := core.NewSessionStore(cfg.Session)
	engine.Use(sessions.Sessions("session", store))

	app.InitRouter(engine)
	engine.HTMLRender = core.LoadTemplates(path.Join(rootpath.RootPath, cfg.TemplateDir), engine.FuncMap)
	//engine.Static("/resource", cfg.ResourceDir)
	engine.StaticFS("/resource", http.Dir(path.Join(rootpath.RootPath, cfg.ResourceDir)))
	fmt.Println(cfg.ServerAddr)
	engine.Run(cfg.ServerAddr)

}
