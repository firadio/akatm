package audit

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuditLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 审计日志详情
func NewGetAuditLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuditLogLogic {
	return &GetAuditLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuditLogLogic) GetAuditLog(req *types.IdReq) (resp *types.GetAuditLogResp, err error) {
	// todo: add your logic here and delete this line

	return
}
