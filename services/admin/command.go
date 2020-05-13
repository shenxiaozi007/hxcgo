package admin

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/huangxinchun/hxcgo/services/admin/app"
	"github.com/huangxinchun/hxcgo/services/admin/core/cache"
	"github.com/huangxinchun/hxcgo/services/admin/core/db"
	"github.com/huangxinchun/hxcgo/services/admin/core/opt"
	"github.com/huangxinchun/hxcgo/services/admin/core/redis"
	"github.com/spf13/cobra"
)

var AdminModel = &cobra.Command{
	Use:   "admin",
	Short: "admin服务",
	Args:  cobra.NoArgs, //没有参数

	Run: func(cmd *cobra.Command, args []string) {
		//获取当前目录
		// pwd, _ := os.Getwd()

		err := opt.ParseConfig("/home/vagrant/gocode/hxcgo/services/admin/conf/config.json")
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

	},
}
