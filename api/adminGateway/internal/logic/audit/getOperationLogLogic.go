package audit

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOperationLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 操作日志详情
func NewGetOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOperationLogLogic {
	return &GetOperationLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOperationLogLogic) GetOperationLog(req *types.IdReq) (resp *types.GetOperationLogResp, err error) {
	// todo: add your logic here and delete this line

	return
}
