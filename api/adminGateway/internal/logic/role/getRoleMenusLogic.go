package role

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/admin/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取角色菜单
func NewGetRoleMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleMenusLogic {
	return &GetRoleMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleMenusLogic) GetRoleMenus(req *types.IdReq) (resp *types.GetRoleMenusResp, err error) {
	// 调用RPC服务获取角色菜单
	rpcReq := &admin.GetRoleMenusReq{
		RoleId:    req.Id,
		RequestId: "req-" + string(rune(req.Id)),
	}

	rpcResp, err := l.svcCtx.AdminRpc.GetRoleMenus(l.ctx, rpcReq)
	if err != nil {
		l.Errorf("调用RPC获取角色菜单失败: %v", err)
		return &types.GetRoleMenusResp{
			Code:    500,
			Message: "获取角色菜单失败",
		}, nil
	}

	// 转换响应格式
	if rpcResp.Base.Code != 200 {
		return &types.GetRoleMenusResp{
			Code:    int32(rpcResp.Base.Code),
			Message: rpcResp.Base.Message,
		}, nil
	}

	// 转换菜单数据
	var menuItems []types.RoleMenuItem
	for _, menu := range rpcResp.Menus {
		menuItems = append(menuItems, types.RoleMenuItem{
			Id:       menu.Id,
			Name:     menu.Name,
			Path:     menu.Path,
			ParentId: menu.ParentId,
			Sort:     int(menu.Sort),
		})
	}

	return &types.GetRoleMenusResp{
		Code:    200,
		Message: "获取成功",
		Data:    menuItems,
	}, nil
}
