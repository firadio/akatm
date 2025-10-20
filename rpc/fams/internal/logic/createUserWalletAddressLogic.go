package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserWalletAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserWalletAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserWalletAddressLogic {
	return &CreateUserWalletAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ============ 钱包地址管理相关 ============
func (l *CreateUserWalletAddressLogic) CreateUserWalletAddress(in *famsRpc.CreateUserWalletAddressReq) (*famsRpc.CreateUserWalletAddressResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.CreateUserWalletAddressResp{}, nil
}
