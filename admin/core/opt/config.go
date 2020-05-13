package opt

import (
	"encoding/json"
	"io/ioutil"

	"github.com/huangxinchun/hxcgo/admin/core"
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

func Config() *core.Config {
	return config
}
