package customer

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuditCustomerKycLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// KYC审核
func NewAuditCustomerKycLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuditCustomerKycLogic {
	return &AuditCustomerKycLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuditCustomerKycLogic) AuditCustomerKyc(req *types.AuditCustomerKycReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
