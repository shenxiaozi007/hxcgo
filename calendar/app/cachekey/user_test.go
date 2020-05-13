package cachekey

import (
	"log"
	"testing"
)

func TestUserByOpenID(t *testing.T) {
	log.Println(UserByOpenID("1000"))
}
