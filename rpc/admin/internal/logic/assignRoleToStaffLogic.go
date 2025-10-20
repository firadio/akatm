package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/orm/table"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignRoleToStaffLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssignRoleToStaffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignRoleToStaffLogic {
	return &AssignRoleToStaffLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AssignRoleToStaff 为员工分配角色
func (l *AssignRoleToStaffLogic) AssignRoleToStaff(in *adminRpc.AssignRoleToStaffReq) (*adminRpc.AssignRoleToStaffResp, error) {
	// 检查员工角色关联是否已存在
	exists, err := l.svcCtx.StaffRoleRepository.ExistsByStaffIDAndRoleID(uint(in.StaffId), uint(in.RoleId))
	if err != nil {
		l.Errorf("检查员工角色关联失败: %v", err)
		return &adminRpc.AssignRoleToStaffResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "系统错误",
				RequestId: in.RequestId,
			},
		}, nil
	}
	if exists {
		return &adminRpc.AssignRoleToStaffResp{
			Base: &adminRpc.BaseResp{
				Code:      400,
				Message:   "员工角色关联已存在",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 创建员工角色关联
	staffRole := &table.StaffRole{
		StaffId: uint(in.StaffId),
		RoleId:  uint(in.RoleId),
	}

	if err := l.svcCtx.StaffRoleRepository.Create(staffRole); err != nil {
		l.Errorf("创建员工角色关联失败: %v", err)
		return &adminRpc.AssignRoleToStaffResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "分配角色失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	return &adminRpc.AssignRoleToStaffResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "分配成功",
			RequestId: in.RequestId,
		},
	}, nil
}
