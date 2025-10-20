package dashboard

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRealtimeMonitorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 实时监控
func NewGetRealtimeMonitorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRealtimeMonitorLogic {
	return &GetRealtimeMonitorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRealtimeMonitorLogic) GetRealtimeMonitor() (resp *types.GetRealtimeMonitorResp, err error) {
	// todo: add your logic here and delete this line

	return
}
