package auditAccountApplication

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountApplicationStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 申请统计
func NewGetAccountApplicationStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountApplicationStatsLogic {
	return &GetAccountApplicationStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAccountApplicationStatsLogic) GetAccountApplicationStats(req *types.GetAccountApplicationStatsReq) (resp *types.GetAccountApplicationStatsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
