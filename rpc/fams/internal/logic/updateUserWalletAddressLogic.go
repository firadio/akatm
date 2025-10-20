package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserWalletAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserWalletAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserWalletAddressLogic {
	return &UpdateUserWalletAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新钱包地址
func (l *UpdateUserWalletAddressLogic) UpdateUserWalletAddress(in *famsRpc.UpdateUserWalletAddressReq) (*famsRpc.UpdateUserWalletAddressResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.UpdateUserWalletAddressResp{}, nil
}
