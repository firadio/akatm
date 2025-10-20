package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductUserWalletLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductUserWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductUserWalletLogic {
	return &DeductUserWalletLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 扣减资金
func (l *DeductUserWalletLogic) DeductUserWallet(in *famsRpc.DeductUserWalletReq) (*famsRpc.DeductUserWalletResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.DeductUserWalletResp{}, nil
}
