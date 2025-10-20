package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRolesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRolesLogic {
	return &ListRolesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListRoles 获取角色列表
func (l *ListRolesLogic) ListRoles(in *adminRpc.ListRolesReq) (*adminRpc.ListRolesResp, error) {
	// 使用 Repository 获取角色列表
	roles, total, err := l.svcCtx.RoleRepository.List(
		in.PageReq.Page,
		in.PageReq.PageSize,
		in.PageReq.Keyword,
	)
	if err != nil {
		l.Errorf("获取角色列表失败: %v", err)
		return &adminRpc.ListRolesResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "获取角色列表失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 转换为响应格式
	var roleInfos []*adminRpc.RoleInfo
	for _, role := range roles {
		roleInfos = append(roleInfos, &adminRpc.RoleInfo{
			Id:          int64(role.ID),
			Name:        role.Name,
			Code:        role.Code,
			Description: role.Description,
			Status:      int32(role.Status),
			CreatedAt:   role.CreatedAt.Unix(),
		})
	}

	return &adminRpc.ListRolesResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "获取成功",
			RequestId: in.RequestId,
		},
		PageResp: &adminRpc.PageResp{
			Total:    total,
			Page:     in.PageReq.Page,
			PageSize: in.PageReq.PageSize,
		},
		Roles: roleInfos,
	}, nil
}
