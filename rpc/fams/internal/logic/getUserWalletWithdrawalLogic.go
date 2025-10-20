package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserWalletWithdrawalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserWalletWithdrawalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserWalletWithdrawalLogic {
	return &GetUserWalletWithdrawalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取提现记录
func (l *GetUserWalletWithdrawalLogic) GetUserWalletWithdrawal(in *famsRpc.GetUserWalletWithdrawalReq) (*famsRpc.GetUserWalletWithdrawalResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.GetUserWalletWithdrawalResp{}, nil
}
