package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMenusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMenusLogic {
	return &ListMenusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取菜单列表
func (l *ListMenusLogic) ListMenus(in *adminRpc.ListMenusReq) (*adminRpc.ListMenusResp, error) {
	// 获取菜单列表
	menus, total, err := l.svcCtx.MenuRepository.List(
		in.PageReq.Page,
		in.PageReq.PageSize,
		in.Keyword,
		int8(in.Type),
	)
	if err != nil {
		l.Errorf("获取菜单列表失败: %v", err)
		return &adminRpc.ListMenusResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "获取菜单列表失败",
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

	return &adminRpc.ListMenusResp{
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
		Menus: menuInfos,
	}, nil
}
