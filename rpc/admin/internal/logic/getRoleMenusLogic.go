package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleMenusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleMenusLogic {
	return &GetRoleMenusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetRoleMenus 获取角色的所有菜单
func (l *GetRoleMenusLogic) GetRoleMenus(in *adminRpc.GetRoleMenusReq) (*adminRpc.GetRoleMenusResp, error) {
	// 获取角色的所有菜单
	menus, err := l.svcCtx.RoleMenuRepository.GetMenusByRoleID(uint(in.RoleId))
	if err != nil {
		l.Errorf("获取角色菜单失败: %v", err)
		return &adminRpc.GetRoleMenusResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "获取角色菜单失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 转换为响应格式
	var menuInfos []*adminRpc.MenuInfo
	for _, menu := range menus {
		menuInfos = append(menuInfos, &adminRpc.MenuInfo{
			Id:        int64(menu.ID),
			Name:      menu.Name,
			Path:      menu.Path,
			Icon:      menu.Icon,
			Component: menu.Component,
			Role:      menu.Role,
			Label:     menu.Label,
			Alias:     menu.Alias,
			Type:      int32(menu.Type),
			ParentId:  int64(menu.ParentId),
			Sort:      int32(menu.Sort),
			Status:    int32(menu.Status),
			CreatedAt: menu.CreatedAt.Unix(),
		})
	}

	return &adminRpc.GetRoleMenusResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "获取成功",
			RequestId: in.RequestId,
		},
		Menus: menuInfos,
	}, nil
}
