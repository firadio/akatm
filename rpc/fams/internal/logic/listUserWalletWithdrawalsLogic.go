package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserWalletWithdrawalsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserWalletWithdrawalsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserWalletWithdrawalsLogic {
	return &ListUserWalletWithdrawalsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取提现记录列表
func (l *ListUserWalletWithdrawalsLogic) ListUserWalletWithdrawals(in *famsRpc.ListUserWalletWithdrawalsReq) (*famsRpc.ListUserWalletWithdrawalsResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.ListUserWalletWithdrawalsResp{}, nil
}
