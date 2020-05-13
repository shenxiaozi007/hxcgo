package uuid

import (
	"crypto/md5"
	"fmt"
	"log"
	"testing"
)

func Test_uuid(t *testing.T) {
	idFactory, _ := New(1)
	log.Println(fmt.Sprintf("%x", md5.Sum([]byte(idFactory.String()))))
}
