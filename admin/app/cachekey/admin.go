package cachekey

func AdminUniqueEmail(email string) string {
	return Sprintf("admin_uni_email:%s", email)
}

func AdminUniqueMobile(mobile string) string {
	return Sprintf("admin_uni_mobile:%s", mobile)
}
