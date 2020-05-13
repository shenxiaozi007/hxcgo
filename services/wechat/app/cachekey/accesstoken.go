package cachekey

func AccessTokenLock(appID string) string {
	return Sprintf("access_token_lock:%s",appID)
}

func AccessToken(appID string) string {
	return Sprintf("access_token:%s",appID)
}