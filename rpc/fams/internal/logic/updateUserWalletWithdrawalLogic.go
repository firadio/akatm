package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserWalletWithdrawalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserWalletWithdrawalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserWalletWithdrawalLogic {
	return &UpdateUserWalletWithdrawalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新提现记录
func (l *UpdateUserWalletWithdrawalLogic) UpdateUserWalletWithdrawal(in *famsRpc.UpdateUserWalletWithdrawalReq) (*famsRpc.UpdateUserWalletWithdrawalResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.UpdateUserWalletWithdrawalResp{}, nil
}
