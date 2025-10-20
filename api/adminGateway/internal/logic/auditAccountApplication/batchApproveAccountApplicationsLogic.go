package auditAccountApplication

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchApproveAccountApplicationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量审核通过
func NewBatchApproveAccountApplicationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchApproveAccountApplicationsLogic {
	return &BatchApproveAccountApplicationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchApproveAccountApplicationsLogic) BatchApproveAccountApplications(req *types.BatchApproveAccountApplicationsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
