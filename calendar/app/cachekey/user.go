package cachekey

func UserByOpenID(openID string) string {
	return Sprintf("user_open_id:%s", openID)
}

func UserWechatSessionKeyByID(id uint) string {
	return Sprintf("user_wx_sess_key:%d", id)
}
