package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新菜单
func (l *UpdateMenuLogic) UpdateMenu(in *adminRpc.UpdateMenuReq) (*adminRpc.UpdateMenuResp, error) {
	// 获取现有菜单
	menu, err := l.svcCtx.MenuRepository.GetByID(uint(in.MenuId))
	if err != nil {
		l.Errorf("获取菜单失败: %v", err)
		return &adminRpc.UpdateMenuResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "获取菜单失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 更新菜单信息
	menu.Name = in.Name
	menu.Path = in.Path
	menu.Icon = in.Icon
	menu.Component = in.Component
	menu.Role = in.Role
	menu.Label = in.Label
	menu.Alias = in.Alias
	menu.Type = int8(in.Type)
	menu.ParentId = uint(in.ParentId)
	menu.Sort = int(in.Sort)
	menu.Status = int8(in.Status)

	if err := l.svcCtx.MenuRepository.Update(menu); err != nil {
		l.Errorf("更新菜单失败: %v", err)
		return &adminRpc.UpdateMenuResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "更新菜单失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 转换为响应格式
	menuInfo := &adminRpc.MenuInfo{
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
	}

	return &adminRpc.UpdateMenuResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "更新成功",
			RequestId: in.RequestId,
		},
		Data: menuInfo,
	}, nil
}
