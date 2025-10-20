package audit

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 日志统计
func NewGetLogStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogStatsLogic {
	return &GetLogStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogStatsLogic) GetLogStats(req *types.GetLogStatsReq) (resp *types.GetLogStatsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
