package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FreezeUserWalletLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFreezeUserWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FreezeUserWalletLogic {
	return &FreezeUserWalletLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ============ 资金操作相关 ============
func (l *FreezeUserWalletLogic) FreezeUserWallet(in *famsRpc.FreezeUserWalletReq) (*famsRpc.FreezeUserWalletResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.FreezeUserWalletResp{}, nil
}
