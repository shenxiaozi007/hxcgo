package opt

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	configPath := "/home/vagrant/gocode/hxcgo/admin/conf/"

	err := ViperConfig(configPath)

	assert.Nilf(t, err, "设置失败")
	log.Println()
}
