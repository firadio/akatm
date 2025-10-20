package audit

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAuditLogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 审计日志列表
func NewListAuditLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAuditLogsLogic {
	return &ListAuditLogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAuditLogsLogic) ListAuditLogs(req *types.ListAuditLogsReq) (resp *types.ListAuditLogsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
