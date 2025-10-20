package dashboard

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuditStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 审核统计
func NewGetAuditStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuditStatsLogic {
	return &GetAuditStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuditStatsLogic) GetAuditStats(req *types.GetAuditStatsReq) (resp *types.GetAuditStatsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
