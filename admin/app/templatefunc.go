package app

import "admin/app/pkg"

func IsGranted(sess *pkg.Session,uri string) bool {
	return sess.IsGranted(uri)
}
