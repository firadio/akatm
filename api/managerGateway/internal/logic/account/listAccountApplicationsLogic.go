package account

import (
	"context"

	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAccountApplicationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 开户申请列表
func NewListAccountApplicationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAccountApplicationsLogic {
	return &ListAccountApplicationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAccountApplicationsLogic) ListAccountApplications(req *types.ListAccountApplicationsReq) (resp *types.ListAccountApplicationsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
