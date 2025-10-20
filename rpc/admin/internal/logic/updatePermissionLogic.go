package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePermissionLogic {
	return &UpdatePermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新权限
func (l *UpdatePermissionLogic) UpdatePermission(in *adminRpc.UpdatePermissionReq) (*adminRpc.UpdatePermissionResp, error) {
	// todo: add your logic here and delete this line

	return &adminRpc.UpdatePermissionResp{}, nil
}
