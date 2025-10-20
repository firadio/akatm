package audit

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportLogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 导出日志
func NewExportLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportLogsLogic {
	return &ExportLogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportLogsLogic) ExportLogs(req *types.ExportLogsReq) (resp *types.ExportLogsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
