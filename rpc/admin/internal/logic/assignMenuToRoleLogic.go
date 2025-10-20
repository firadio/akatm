package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/orm/table"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignMenuToRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssignMenuToRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignMenuToRoleLogic {
	return &AssignMenuToRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AssignMenuToRole 为角色分配菜单
func (l *AssignMenuToRoleLogic) AssignMenuToRole(in *adminRpc.AssignMenuToRoleReq) (*adminRpc.AssignMenuToRoleResp, error) {
	// 检查角色菜单关联是否已存在
	exists, err := l.svcCtx.RoleMenuRepository.ExistsByRoleIDAndMenuID(uint(in.RoleId), uint(in.MenuId))
	if err != nil {
		l.Errorf("检查角色菜单关联失败: %v", err)
		return &adminRpc.AssignMenuToRoleResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "系统错误",
				RequestId: in.RequestId,
			},
		}, nil
	}
	if exists {
		return &adminRpc.AssignMenuToRoleResp{
			Base: &adminRpc.BaseResp{
				Code:      400,
				Message:   "角色菜单关联已存在",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 创建角色菜单关联
	roleMenu := &table.RoleMenu{
		RoleId: uint(in.RoleId),
		MenuId: uint(in.MenuId),
	}

	if err := l.svcCtx.RoleMenuRepository.Create(roleMenu); err != nil {
		l.Errorf("创建角色菜单关联失败: %v", err)
		return &adminRpc.AssignMenuToRoleResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "分配菜单失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	return &adminRpc.AssignMenuToRoleResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "分配成功",
			RequestId: in.RequestId,
		},
	}, nil
}
