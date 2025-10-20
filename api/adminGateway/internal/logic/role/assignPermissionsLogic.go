package role

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignPermissionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分配权限
func NewAssignPermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignPermissionsLogic {
	return &AssignPermissionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignPermissionsLogic) AssignPermissions(req *types.AssignPermissionsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
