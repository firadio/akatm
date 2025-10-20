package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnfreezeUserWalletLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnfreezeUserWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfreezeUserWalletLogic {
	return &UnfreezeUserWalletLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 解冻资金
func (l *UnfreezeUserWalletLogic) UnfreezeUserWallet(in *famsRpc.UnfreezeUserWalletReq) (*famsRpc.UnfreezeUserWalletResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.UnfreezeUserWalletResp{}, nil
}
