package dashboard

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTrendChartsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 趋势图表
func NewGetTrendChartsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTrendChartsLogic {
	return &GetTrendChartsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTrendChartsLogic) GetTrendCharts(req *types.GetTrendChartsReq) (resp *types.GetTrendChartsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
