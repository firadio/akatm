package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePermissionLogic {
	return &CreatePermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ============ 权限管理相关 ============
func (l *CreatePermissionLogic) CreatePermission(in *adminRpc.CreatePermissionReq) (*adminRpc.CreatePermissionResp, error) {
	// todo: add your logic here and delete this line

	return &adminRpc.CreatePermissionResp{}, nil
}
