package dashboard

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCustomerStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 客户统计
func NewGetCustomerStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCustomerStatsLogic {
	return &GetCustomerStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCustomerStatsLogic) GetCustomerStats(req *types.GetCustomerStatsReq) (resp *types.GetCustomerStatsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
