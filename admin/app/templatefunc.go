package app

import "github.com/huangxinchun/hxcgo/admin/app/pkg"

func IsGranted(sess *pkg.Session,uri string) bool {
	return sess.IsGranted(uri)
}
