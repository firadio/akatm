package auditWithdrawal

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchApproveWithdrawalsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量审核通过
func NewBatchApproveWithdrawalsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchApproveWithdrawalsLogic {
	return &BatchApproveWithdrawalsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchApproveWithdrawalsLogic) BatchApproveWithdrawals(req *types.BatchApproveWithdrawalsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
