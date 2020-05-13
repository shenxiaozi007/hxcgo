package ctrl

import "wechat/core"
import "wechat/core/opt"
import (
	e "wechat/app/err"
)

func getWechatConfig(appName string) (*core.WechatConfig,error) {
	if appName == "" {
		return nil,e.EParam
	}
	cfg := opt.Config()

	wxConfig, ok := cfg.Wechat[appName]
	if !ok {
		return nil,e.EInvalidAppName
	}

	return wxConfig,nil
}
