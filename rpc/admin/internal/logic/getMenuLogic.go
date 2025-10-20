package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuLogic {
	return &GetMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取菜单
func (l *GetMenuLogic) GetMenu(in *adminRpc.GetMenuReq) (*adminRpc.GetMenuResp, error) {
	// 获取菜单
	menu, err := l.svcCtx.MenuRepository.GetByID(uint(in.MenuId))
	if err != nil {
		l.Errorf("获取菜单失败: %v", err)
		return &adminRpc.GetMenuResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "获取菜单失败",
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

	return &adminRpc.GetMenuResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "获取成功",
			RequestId: in.RequestId,
		},
		Data: menuInfo,
	}, nil
}
