package auditWithdrawal

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWithdrawalStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 提现统计
func NewGetWithdrawalStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWithdrawalStatsLogic {
	return &GetWithdrawalStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWithdrawalStatsLogic) GetWithdrawalStats(req *types.GetWithdrawalStatsReq) (resp *types.GetWithdrawalStatsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
