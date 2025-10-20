package auditDeposit

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApproveDepositLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 审核通过
func NewApproveDepositLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApproveDepositLogic {
	return &ApproveDepositLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApproveDepositLogic) ApproveDeposit(req *types.ApproveDepositReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
