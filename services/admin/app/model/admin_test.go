package model

import (
	"github.com/huangxinchun/hxcgo/services/admin/core"
	"github.com/huangxinchun/hxcgo/services/admin/core/db"
	"log"
	"testing"
)

func TestAdmin(t *testing.T)  {
	dbConfigs := []*core.DBConfig{
		&core.DBConfig{
			Driver:"mysql",
			Alias:"default",
			Username:"root",
			Password:"123456",
			Host:"localhost",
			Port:3306,
			Database:"faker",
			TablePrefix:"f_",
		},
	}
	log.Println(db.Connect(dbConfigs))

	//loginAt := time.Now()
	//admin := &Admin{
	//	Name:"dylan",
	//	Email:"123@qq.com",
	//	Mobile:"13650515454",
	//	Password:"tttt",
	//	LoginAt:&loginAt,
	//}
	//
	//db := core.GetDB().Create(admin)

	//adminModel := Admin{}
	//admin,isExists := adminModel.GetAdmin(2)
	//
	//log.Printf("%#v",admin)
	//log.Println(admin.CreatedAt,admin.UpdatedAt)
	//log.Println(isExists)
	//log.Println(core.GetDB().GetErrors())

}

