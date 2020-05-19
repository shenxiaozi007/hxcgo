package opt

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	configPath := "/home/vagrant/gocode/hxcgo/admin/conf/"

	err := ViperConfig(configPath)

	fmt.Printf("p1=%#v", config)
	assert.Nilf(t, err, "设置失败")
	log.Println()
}
