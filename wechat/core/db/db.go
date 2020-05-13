package db

import (
	"fmt"
	"wechat/core"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbs = map[string]*gorm.DB{}

func Connect(configs []*core.DBConfig) error {
	var defaultDB *gorm.DB

	var err error
	var db *gorm.DB
	for _, cfg := range configs {
		switch cfg.Driver {
		case "mysql":
			db, err = ConnectMysql(cfg)
			if err != nil {
				return err
			}

			dbs[cfg.Alias] = db
		default:
			return fmt.Errorf("driver %s not exists", cfg.Driver)
		}

		if defaultDB == nil || cfg.Alias == "default" {
			dbs["default"] = db
		}
	}

	return nil
}

func ConnectMysql(cfg *core.DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database))
	if err != nil {
		return nil, fmt.Errorf("can not connect database")
	}
	//db.SingularTable(true)
	db.LogMode(true) //debug
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return cfg.TablePrefix + defaultTableName
	}

	return db, nil
}

func Client(alias ...string) *gorm.DB {
	var dbName string
	if len(alias) == 0 {
		dbName = "default"
	} else {
		dbName = alias[0]
	}

	db, ok := dbs[dbName]
	if !ok {
		panic(fmt.Sprintf("database %s does not exists", dbName))
	}

	return db
}
