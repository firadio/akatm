package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserWalletAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserWalletAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserWalletAddressLogic {
	return &DeleteUserWalletAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除钱包地址
func (l *DeleteUserWalletAddressLogic) DeleteUserWalletAddress(in *famsRpc.DeleteUserWalletAddressReq) (*famsRpc.DeleteUserWalletAddressResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.DeleteUserWalletAddressResp{}, nil
}
