package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserWalletLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserWalletLogic {
	return &GetUserWalletLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户钱包
func (l *GetUserWalletLogic) GetUserWallet(in *famsRpc.GetUserWalletReq) (*famsRpc.GetUserWalletResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.GetUserWalletResp{}, nil
}
