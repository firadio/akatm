package auditWithdrawal

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApproveWithdrawalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 审核通过
func NewApproveWithdrawalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApproveWithdrawalLogic {
	return &ApproveWithdrawalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApproveWithdrawalLogic) ApproveWithdrawal(req *types.ApproveWithdrawalReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
