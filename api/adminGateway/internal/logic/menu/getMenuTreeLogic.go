package menu

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 菜单树
func NewGetMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuTreeLogic {
	return &GetMenuTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuTreeLogic) GetMenuTree(req *types.ListMenuReq) (resp *types.ListMenuResp, err error) {
	// 调用RPC服务获取菜单列表
	rpcResp, err := l.svcCtx.AdminRpc.ListMenus(l.ctx, &adminRpc.ListMenusReq{
		PageReq: &adminRpc.PageReq{
			Page:     1,
			PageSize: 1000, // 获取所有菜单用于树形结构
		},
		Keyword:   req.Keyword,
		Type:      1, // 只获取菜单类型，不包含按钮
		RequestId: "get-menu-tree",
	})
	if err != nil {
		l.Errorf("调用RPC获取菜单树失败: %v", err)
		return &types.ListMenuResp{
			Code:    500,
			Message: "获取菜单树失败",
		}, nil
	}

	// 转换响应格式
	var menuInfos []types.MenuInfo
	for _, menu := range rpcResp.Menus {
		menuInfos = append(menuInfos, types.MenuInfo{
			Id:        menu.Id,
			Name:      menu.Name,
			Path:      menu.Path,
			Icon:      menu.Icon,
			Component: menu.Component,
			Role:      menu.Role,
			Label:     menu.Label,
			Alias:     menu.Alias,
			Type:      int(menu.Type),
			ParentId:  menu.ParentId,
			Sort:      int(menu.Sort),
		})
	}

	return &types.ListMenuResp{
		Code:    int32(rpcResp.Base.Code),
		Message: rpcResp.Base.Message,
		Data:    menuInfos,
	}, nil
}
