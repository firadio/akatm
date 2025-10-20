package audit

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExportTaskStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取导出任务状态
func NewGetExportTaskStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExportTaskStatusLogic {
	return &GetExportTaskStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExportTaskStatusLogic) GetExportTaskStatus(req *types.GetExportTaskStatusReq) (resp *types.GetExportTaskStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
