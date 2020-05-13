package cachekey

import "fmt"

const prefix = "sadmin_"

func Sprintf(format string, v ...interface{}) string {
	return fmt.Sprintf(prefix+format, v...)
}
