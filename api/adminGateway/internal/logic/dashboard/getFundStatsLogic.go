package dashboard

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFundStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 资金统计
func NewGetFundStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFundStatsLogic {
	return &GetFundStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFundStatsLogic) GetFundStats(req *types.GetFundStatsReq) (resp *types.GetFundStatsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
