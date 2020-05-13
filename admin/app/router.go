package app

import (
	"github.com/huangxinchun/hxcgo/admin/app/ctrl"
	"net/rpc"
)

func InitRouter() error {
	rpc.Register(ctrl.NewAdmin())
	rpc.Register(ctrl.NewGroup())
	rpc.Register(ctrl.NewRole())
	rpc.Register(ctrl.NewPrivilege())

	rpc.HandleHTTP()
	return nil
}
