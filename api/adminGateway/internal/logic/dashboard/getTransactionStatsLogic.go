package dashboard

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTransactionStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 交易统计
func NewGetTransactionStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTransactionStatsLogic {
	return &GetTransactionStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTransactionStatsLogic) GetTransactionStats(req *types.GetTransactionStatsReq) (resp *types.GetTransactionStatsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
