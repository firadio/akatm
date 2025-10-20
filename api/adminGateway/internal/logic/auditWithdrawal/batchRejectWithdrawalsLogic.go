package auditWithdrawal

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchRejectWithdrawalsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量审核拒绝
func NewBatchRejectWithdrawalsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchRejectWithdrawalsLogic {
	return &BatchRejectWithdrawalsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchRejectWithdrawalsLogic) BatchRejectWithdrawals(req *types.BatchRejectWithdrawalsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
