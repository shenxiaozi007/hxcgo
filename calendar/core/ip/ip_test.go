package ip

import (
	"log"
	"net"
	"testing"
)

func TestToInt64(t *testing.T) {
	ipAddr := net.ParseIP("127.0.0.1")

	long := ToInt64(ipAddr)
	log.Println("127.0.0.1 int64= ", long)

	log.Println(ToIP(long).String())

	ipAddr = net.ParseIP("22.215.255.5")

	long = ToInt64(ipAddr)
	log.Println("22.215.255.5 int64= ", long)

	log.Println(ToIP(long).String())
}
