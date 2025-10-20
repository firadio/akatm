package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/orm/table"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuLogic {
	return &CreateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateMenu 创建菜单
func (l *CreateMenuLogic) CreateMenu(in *adminRpc.CreateMenuReq) (*adminRpc.CreateMenuResp, error) {
	// 创建菜单
	menu := &table.Menu{
		Name:      in.Name,
		Path:      in.Path,
		Icon:      in.Icon,
		Component: in.Component,
		Role:      in.Role,
		Label:     in.Label,
		Alias:     in.Alias,
		Type:      int8(in.Type),
		ParentId:  uint(in.ParentId),
		Sort:      int(in.Sort),
		Status:    1, // 默认启用
	}

	if err := l.svcCtx.MenuRepository.Create(menu); err != nil {
		l.Errorf("创建菜单失败: %v", err)
		return &adminRpc.CreateMenuResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "创建菜单失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 返回创建的菜单信息
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

	return &adminRpc.CreateMenuResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "创建成功",
			RequestId: in.RequestId,
		},
		Data: menuInfo,
	}, nil
}
