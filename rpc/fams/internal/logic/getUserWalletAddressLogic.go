package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserWalletAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserWalletAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserWalletAddressLogic {
	return &GetUserWalletAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取钱包地址
func (l *GetUserWalletAddressLogic) GetUserWalletAddress(in *famsRpc.GetUserWalletAddressReq) (*famsRpc.GetUserWalletAddressResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.GetUserWalletAddressResp{}, nil
}
