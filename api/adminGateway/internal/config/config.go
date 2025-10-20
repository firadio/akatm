package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Signature struct {
		Enable      bool
		Salt        string
		SkewSeconds int64
	}
	AdminRpc zrpc.RpcClientConf
	FamsRpc  zrpc.RpcClientConf
	IamRpc   zrpc.RpcClientConf
}
