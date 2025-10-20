package svc

import (
	"akatm/api/adminGateway/internal/config"
	"akatm/rpc/admin/admin"
	"akatm/rpc/fams/fams"
	"akatm/rpc/iam/iam"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	SignCheck rest.Middleware
	JwtAuth   rest.Middleware
	AdminRpc  admin.Admin
	FamsRpc   fams.Fams
	IamRpc    iam.Iam
}

func NewServiceContext(c config.Config) *ServiceContext {
	// RPC客户端
	adminRpcClient := zrpc.MustNewClient(c.AdminRpc)
	adminRpc := admin.NewAdmin(adminRpcClient)

	famsRpcClient := zrpc.MustNewClient(c.FamsRpc)
	famsRpc := fams.NewFams(famsRpcClient)

	iamRpcClient := zrpc.MustNewClient(c.IamRpc)
	iamRpc := iam.NewIam(iamRpcClient)

	svcCtx := &ServiceContext{
		Config:   c,
		AdminRpc: adminRpc,
		FamsRpc:  famsRpc,
		IamRpc:   iamRpc,
	}

	return svcCtx
}
