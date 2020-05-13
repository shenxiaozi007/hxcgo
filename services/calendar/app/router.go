package app

import (
	"calendar/app/ctrl"
	"net/rpc"
)

func InitRouter() error {
	rpc.Register(ctrl.NewUser())

	rpc.HandleHTTP()
	return nil
}
