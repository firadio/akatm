package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuditBankAccountApplicationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuditBankAccountApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuditBankAccountApplicationLogic {
	return &AuditBankAccountApplicationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 审核银行账户申请
func (l *AuditBankAccountApplicationLogic) AuditBankAccountApplication(in *famsRpc.AuditBankAccountApplicationReq) (*famsRpc.AuditBankAccountApplicationResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.AuditBankAccountApplicationResp{}, nil
}
