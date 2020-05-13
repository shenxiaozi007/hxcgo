package hxcgo

import (
	"github.com/huangxinchun/hxcgo/admin/core/redis"
	"log"
)

//test
func main() {
	err := redis.Connect(cfg.Redis)
	if err != nil {
		log.Fatalln("redis connect error: ", err)
	}
}