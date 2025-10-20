package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuditUserWalletWithdrawalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuditUserWalletWithdrawalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuditUserWalletWithdrawalLogic {
	return &AuditUserWalletWithdrawalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 审核提现申请
func (l *AuditUserWalletWithdrawalLogic) AuditUserWalletWithdrawal(in *famsRpc.AuditUserWalletWithdrawalReq) (*famsRpc.AuditUserWalletWithdrawalResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.AuditUserWalletWithdrawalResp{}, nil
}
