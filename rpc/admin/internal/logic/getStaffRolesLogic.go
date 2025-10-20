package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStaffRolesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStaffRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStaffRolesLogic {
	return &GetStaffRolesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetStaffRoles 获取员工的所有角色
func (l *GetStaffRolesLogic) GetStaffRoles(in *adminRpc.GetStaffRolesReq) (*adminRpc.GetStaffRolesResp, error) {
	// 获取员工的所有角色
	roles, err := l.svcCtx.StaffRoleRepository.GetRolesByStaffID(uint(in.StaffId))
	if err != nil {
		l.Errorf("获取员工角色失败: %v", err)
		return &adminRpc.GetStaffRolesResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "获取员工角色失败",
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

	return &adminRpc.GetStaffRolesResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "获取成功",
			RequestId: in.RequestId,
		},
		Roles: roleInfos,
	}, nil
}
