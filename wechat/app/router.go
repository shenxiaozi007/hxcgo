package app

import (
	"net/rpc"
	"wechat/app/ctrl"
)

func InitRouter() error {
	rpc.Register(ctrl.NewMiniProgram())
	rpc.Register(ctrl.NewWechat())
	rpc.Register(ctrl.NewSession())
	rpc.Register(ctrl.NewQRCode())
	rpc.HandleHTTP()
	return nil
}
