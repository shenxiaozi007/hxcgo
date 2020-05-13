package cachekey

import "fmt"

const prefix = "nob_"

func Sprintf(format string, v ...interface{}) string {
	return fmt.Sprintf(prefix+format, v...)
}
