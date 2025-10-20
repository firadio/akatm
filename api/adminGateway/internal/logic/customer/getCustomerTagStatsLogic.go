package customer

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCustomerTagStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取标签使用统计
func NewGetCustomerTagStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCustomerTagStatsLogic {
	return &GetCustomerTagStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCustomerTagStatsLogic) GetCustomerTagStats(req *types.GetCustomerTagStatsReq) (resp *types.GetCustomerTagStatsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
