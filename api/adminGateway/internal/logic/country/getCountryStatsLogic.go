package country

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCountryStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 国家使用统计
func NewGetCountryStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCountryStatsLogic {
	return &GetCountryStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCountryStatsLogic) GetCountryStats(req *types.GetCountryStatsReq) (resp *types.GetCountryStatsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
