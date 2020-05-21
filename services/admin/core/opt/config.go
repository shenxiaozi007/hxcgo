package opt

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/huangxinchun/hxcgo/services/admin/core"
	"github.com/spf13/viper"
)

var config = &core.Config{}

//ParseConfig
func ParseConfig(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, config)
	return err
}

func ViperConfig(filename string) error {

	viperConfig := viper.New()

	viperConfig.AddConfigPath(filename)
	viperConfig.SetConfigType("json")
	viperConfig.SetConfigName("config")

	if err := viperConfig.ReadInConfig(); err != nil {
		panic("读取配置文件错误。" + err.Error())
	}

	viperConfig.WatchConfig()
	viperConfig.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("配置文件已修改成功.修改的配置文件名为:%s", in.Name)
	})

	//json转结构体
	err := viperConfig.Unmarshal(config)
	if err != nil {
		panic("配置文件序列化有误。" + err.Error())
	}

	return err
}

func Config() *core.Config {
	return config
}
