package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserWalletLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserWalletLogic {
	return &UpdateUserWalletLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户钱包
func (l *UpdateUserWalletLogic) UpdateUserWallet(in *famsRpc.UpdateUserWalletReq) (*famsRpc.UpdateUserWalletResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.UpdateUserWalletResp{}, nil
}
