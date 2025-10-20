package audit

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CleanLogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 清理日志
func NewCleanLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanLogsLogic {
	return &CleanLogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CleanLogsLogic) CleanLogs(req *types.CleanLogsReq) (resp *types.CleanLogsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
