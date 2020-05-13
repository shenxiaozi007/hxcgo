package opt

import (
	"admin/core"
	"encoding/json"
	"io/ioutil"
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
