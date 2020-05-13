package dao

import (
	"github.com/huangxinchun/hxcgo/admin/core/db"
	"github.com/huangxinchun/hxcgo/admin/core/opt"
	"log"
	"testing"
)

func connectDB() {
	err := opt.ParseConfig("../../conf/config.json")
	if err != nil {
		log.Fatalln("parser config file err: ", err)
	}

	log.Println(db.Connect(opt.Config().DB))
}

func TestAdminRole_List(t *testing.T) {
	connectDB()

	roleDAO := &Role{}
	roles, err := roleDAO.Roles(nil, 1, 10)
	if err != nil {
		log.Fatalln("Role.List error: ", err)
	}

	for _, role := range roles {
		log.Printf("Role.List: %#v", role)
	}
}

func TestRole_Count(t *testing.T) {
	connectDB()

	roleDAO := &Role{}
	count, err := roleDAO.Count(nil)
	if err != nil {
		log.Fatalln("Role.Count error: ", err)
	}
	log.Printf("Role.Count: %d", count)
}
