package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPermissionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPermissionsLogic {
	return &ListPermissionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取权限列表
func (l *ListPermissionsLogic) ListPermissions(in *adminRpc.ListPermissionsReq) (*adminRpc.ListPermissionsResp, error) {
	// todo: add your logic here and delete this line

	return &adminRpc.ListPermissionsResp{}, nil
}
