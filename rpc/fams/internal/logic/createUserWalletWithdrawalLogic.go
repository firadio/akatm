package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserWalletWithdrawalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserWalletWithdrawalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserWalletWithdrawalLogic {
	return &CreateUserWalletWithdrawalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ============ 提现管理相关 ============
func (l *CreateUserWalletWithdrawalLogic) CreateUserWalletWithdrawal(in *famsRpc.CreateUserWalletWithdrawalReq) (*famsRpc.CreateUserWalletWithdrawalResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.CreateUserWalletWithdrawalResp{}, nil
}
