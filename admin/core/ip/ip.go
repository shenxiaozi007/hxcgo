package ip

import (
	"net"
	"strconv"
	"strings"
)

func ToInt64(ip net.IP) int64 {
	arr := strings.Split(ip.String(), ".")
	if len(arr) != 4 {
		return 0
	}

	p0, _ := strconv.Atoi(arr[0])
	p1, _ := strconv.Atoi(arr[1])
	p2, _ := strconv.Atoi(arr[2])
	p3, _ := strconv.Atoi(arr[3])

	var long int64
	long += int64(p0) << 24
	long += int64(p1) << 16
	long += int64(p2) << 8
	long += int64(p3)
	return long
}

func ToIP(long int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(long & 0xFF)
	bytes[1] = byte((long >> 8) & 0xFF)
	bytes[2] = byte((long >> 16) & 0xFF)
	bytes[3] = byte((long >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}
