package rpc

import (
	"github.com/huangxinchun/hxcgo/admin/core"
	"fmt"
	"net/rpc"
	"time"
)

type ServiceClient struct {
	*rpc.Client
	cfg *core.RPCConfig
	shutdown bool
}


var clients = map[string]*ServiceClient{}

func Connect(configs []*core.RPCConfig) error {
	for _,cfg := range configs {
		client,err := rpc.DialHTTP("tcp",cfg.Addr)
		if err != nil {
			return err
		}

		clients[cfg.ServiceName] = &ServiceClient{
			Client:   client,
			shutdown: false,
			cfg:cfg,
		}
	}

	go func() {
		ticker := time.Tick(time.Second * 3)
		for{
			select {
			case <-ticker:
				for _,client := range clients {
					if client.shutdown {
						cli,err := rpc.DialHTTP("tcp",client.cfg.Addr)
						if err != nil {
							continue
						}

						client.Client = cli
						client.shutdown = false
					}
				}
			}
		}
	}()

	return nil
}

func Service(name string) *ServiceClient {
	client,ok := clients[name]
	if !ok {
		panic(fmt.Sprintf("rpc service %s does not exists",name))
	}

	return client
}

func(s *ServiceClient) HandleErr(err error) {
	if err == rpc.ErrShutdown {
		s.shutdown = true
	}
}

func(s *ServiceClient) Call(serviceMethod string, args interface{}, reply interface{}) error {
	err := s.Client.Call(serviceMethod,args,reply)
	s.HandleErr(err)

	return err
}

