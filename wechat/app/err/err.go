package err

import "errors"

var (
	EParam          = errors.New("param error")
	ENotFound       = errors.New("not found")
	EInvalidCode    = errors.New("param error: invalid code")
	EAPILimit       = errors.New("api error: rate limit")
	ESystemBusy     = errors.New("api error: system busy")
	EUnknown        = errors.New("api error: unknown error")
	EInvalidAppID   = errors.New("param error: invalid app id")
	EInvalidAppName = errors.New("param error: invalid app name")
	EInvalidAppSecret = errors.New("invalid app secret")
	EInvalidGrantType = errors.New("invalid grant type")
	EInvalidAccessToken = errors.New("invalid access token")
	EPageNotFoundOrMiniProgramNotPublish = errors.New("page not found or miniprogram not publish")
)
