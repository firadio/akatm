package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserWalletAddressesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserWalletAddressesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserWalletAddressesLogic {
	return &ListUserWalletAddressesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取钱包地址列表
func (l *ListUserWalletAddressesLogic) ListUserWalletAddresses(in *famsRpc.ListUserWalletAddressesReq) (*famsRpc.ListUserWalletAddressesResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.ListUserWalletAddressesResp{}, nil
}
